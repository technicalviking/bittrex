package bittrex

// placeholder
import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/thebotguys/signalr"
)

type clientMethod = func(string, string, []json.RawMessage)

type bittrexSubParseError struct {
	hub        string
	method     string
	message    json.RawMessage
	rawParse   string
	parseError error
	//
}

func (b bittrexSubParseError) Error() string {
	return fmt.Sprintf(
		"Hub: %s \nMethod: %s \nMessage: %s \nRaw Parse: %v \nOriginal Error: %s",
		b.hub,
		b.method,
		b.message,
		b.rawParse,
		b.parseError.Error(),
	)
}

//BittrexSubscription struct representing a connection to the Bittrex websocket API.
//sending a value to Done or closing Done
type BittrexSubscription struct {
	Data     chan ExchangeState
	Error    chan error
	Done     chan bool
	market   string
	wsClient *signalr.Client
	timeout  time.Duration
}

func newSub(market string, timeout time.Duration) *BittrexSubscription {
	return &BittrexSubscription{
		make(chan ExchangeState),
		make(chan error),
		make(chan bool),
		market,
		signalr.NewWebsocketClient(),
		timeout,
	}
}

func (b *BittrexSubscription) setSubClientMethods(filter string) {
	b.wsClient.OnClientMethod = func(hub string, method string, msgs []json.RawMessage) {
		if hub != websocketHub || method != filter {
			return
		}
		for _, msg := range msgs {
			b.parseMessage(hub, method, msg)
		}

	}

	b.wsClient.OnMessageError = func(err error) {
		b.Error <- fmt.Errorf("Remote Error: %s", err.Error())
	}
}

func (b *BittrexSubscription) parseMessage(hub, method string, msg json.RawMessage) {

	var exchangeState ExchangeState

	if parseErr := json.Unmarshal(msg, &exchangeState); parseErr != nil {
		b.Error <- newSubError(hub, method, msg, parseErr)
		return
	}

	if b.market == "" || exchangeState.MarketName == b.market {
		b.Data <- exchangeState
	}
}

func (b *BittrexSubscription) connect() {

	connectDone := make(chan error)

	go func() {
		u := url.URL{Scheme: "https", Host: websocketBaseURI, Path: "/signalr/negotiate"}
		fmt.Println(u.String())

		connectDone <- b.wsClient.Connect("https", websocketBaseURI, []string{websocketHub})
	}()

	select {
	case <-time.After(b.timeout):
		b.Error <- fmt.Errorf("timeout error")
	case connectErr := <-connectDone:
		b.Error <- fmt.Errorf("connection error: %v", connectErr)
	}
}

func (b *BittrexSubscription) subToMarket() {
	subMethod := "SubscribeToExchangeDeltas"
	queryMethod := "QueryExchangeState" //return

	var callHubErr error
	var param interface{}

	if b.market != "" {
		param = b.market
	}

	if _, callHubErr = b.wsClient.CallHub(websocketHub, subMethod, param); callHubErr != nil {
		b.Error <- fmt.Errorf("SubToMarket Error: %s", callHubErr.Error())
		//b.wsClient.Close()
		return
	}

	var queryResponse json.RawMessage

	if queryResponse, callHubErr = b.wsClient.CallHub(websocketHub, queryMethod, param); callHubErr != nil {
		b.Error <- fmt.Errorf("QueryExchangeState Error: %s", callHubErr.Error())
		b.wsClient.Close()
		return
	}

	b.parseMessage(websocketHub, queryMethod, queryResponse)
}

func newSubError(hub, method string, msg json.RawMessage, parseErr error) error {
	var rawParse string
	//try again
	if debugParseError := json.Unmarshal(msg, &rawParse); debugParseError != nil {
		rawParse = fmt.Sprintf("Couldn't even parse as string: %s", debugParseError.Error())
	}

	return bittrexSubParseError{
		hub,
		method,
		msg,
		rawParse,
		parseErr,
	}
}

//WsSubExchangeUpdates - Undocumented websocket endpoint for bittrex
//market is an optional parameter.  passing an empty string subscribes to changes for all markets.
//(actually that happens anyway, the param just filters what gets sent to the returned chan)
//***DO NOT USE!  Have to figure out a way around the cloudflare DDOS protection, or wait
//until bittrex deploys an official documented websocket API.***
func (c *Client) WsSubExchangeUpdates(market string) *BittrexSubscription {
	sub := newSub(market, c.timeout)

	sub.setSubClientMethods("updateExchangeState")

	go func() {
		sub.connect()
		sub.subToMarket()

		select {
		case <-sub.Done:
			sub.wsClient.Close()
			close(sub.Data)
			close(sub.Error)
		case <-sub.wsClient.DisconnectedChannel:
			sub.Error <- fmt.Errorf("socket closed by remote host")
			close(sub.Data)
			close(sub.Error)
		}
	}()

	return sub
}
