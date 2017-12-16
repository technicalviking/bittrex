package bittrex

import "net/http"

const (
	baseURI                string = "https://bittrex.com/api"
	apiVersion             string = "v1.1"
	undocumentedApiVersion string = "v2.0"
)

var (
	httpClient = &http.Client{}
)

//Client ...
type Client struct {
	apiKey    string
	apiSecret string
	err       *bittrexError
}

//NewClient initialize the library with a key/secret pair.
func New(key string, secret string) *Client {
	return &Client{
		apiKey:    key,
		apiSecret: secret,
		err:       &bittrexError{},
	}
}

func init() {}
