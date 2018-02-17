package bittrex

import "net/http"
import "time"

const (
	baseURI                string = "https://bittrex.com/api"
	apiVersion             string = "v1.1"
	undocumentedAPIVersion string = "v2.0"
	websocketBaseURI       string = "socket.bittrex.com"
	websocketHub           string = "CoreHub" //SignalR main hub
	defaultTimeout         int64  = 30
)

var (
	httpClient = &http.Client{}
)

//Client ...
type Client struct {
	apiKey    string
	apiSecret string
	err       *bittrexError
	timeout   time.Duration
}

//New initialize the library with a key/secret pair.
func New(key string, secret string) *Client {
	return NewWithCustomTimeout(key, secret, defaultTimeout)
}

//NewWithCustomTimeout initialize the library with a key/secret pair and a custom timeout.
func NewWithCustomTimeout(key string, secret string, seconds int64) *Client {
	return &Client{
		apiKey:    key,
		apiSecret: secret,
		err:       nil,
		timeout:   time.Duration(seconds) * time.Second,
	}
}

func init() {}
