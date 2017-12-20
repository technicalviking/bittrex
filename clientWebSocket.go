package bittrex

// placeholder
import (
	"encoding/json"
	"time"

	"github.com/thebotguys/signalr"

	_ "github.com/thebotguys/signalr"
)

//WsSubExchangeUpdates - Undocumented websocket endpoint for bittrex
//market is an optional parameter.  passing an empty string subscribes to changes for all markets.
//(actually that happens anyway, the param just filters what gets sent to the returned chan)
func (c *Client) WsSubExchangeUpdates(kill <-chan bool, market string) (chan<- ExchangeState, error) {
	wsClient := signalr.NewWebsocketClient()

	dataChan := make(chan<- ExchangeState)

	//what does this do?
	wsClient.OnClientMethod = getOnClientMethod("updateExchangeState", dataChan)

	return nil, nil
}

type onClientMethod = func(string, string, []json.RawMessage)

func getOnClientMethod(filter string, dataChan chan<- ExchangeState, market) onClientMethod {
	return func(hub string, method string, msgs []json.RawMessage) {
		if hub != websocketHub || method != filter {
			return
		}

		for _, msg := range msgs {
			var exchangeState ExchangeState

			if parseErr := json.Unmarshal(msg, &exchangeState); parseErr != nil {
				debugOutput(hub, method, msg)
			}

			if market == "" || exchangeState.MarketName == market {
				dataChan <- exchangeState
			}

			
		}
	}
}

func debugOutput(hub, method string, msg json.RawMessage) {
	var rawParse string
	//try again
	if debugParseError := json.Unmarshal(msg, &rawParse); debugParseError != nil {
		rawParse = "Couldn't even parse as string"
	}
	//fmt.Printf("Hub: %s \nMethod: %s \nMessage: %s \nRawMessage: %v \n", hub, method, rawParse, msg)
}

// SubscribeExchangeUpdate subscribes for updates of the market.
// Updates will be sent to dataCh.
// To stop subscription, send to, or close 'stop'.
func (b *Client) SubscribeExchangeUpdate(market string, dataCh chan<- ExchangeState, stop <-chan bool) error {
	const timeout = 5 * time.Second
	client := signalr.NewWebsocketClient()
	client.OnClientMethod = func(hub string, method string, messages []json.RawMessage) {
		if hub != WS_HUB || method != "updateExchangeState" {
			return
		}
		parseStates(messages, dataCh, market)
	}
	err := doAsyncTimeout(func() error {
		return client.Connect("https", WS_BASE, []string{WS_HUB})
	}, func(err error) {
		if err == nil {
			client.Close()
		}
	}, timeout)
	if err != nil {
		return err
	}
	defer client.Close()
	var msg json.RawMessage
	err = doAsyncTimeout(func() error {
		var err error
		msg, err = subForMarket(client, market)
		return err
	}, nil, timeout)
	if err != nil {
		return err
	}
	var st ExchangeState
	if err = json.Unmarshal(msg, &st); err != nil {
		return err
	}
	st.Initial = true
	st.MarketName = market
	sendStateAsync(dataCh, st)
	select {
	case <-stop:
	case <-client.DisconnectedChannel:
	}
	return nil
}
