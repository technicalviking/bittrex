package bittrex

import "testing"

func TestGetFullURI(t *testing.T) {
	endpoint := "pub/market/getticks"

	params := map[string]string{
		"marketName":   "BTC-LTC",
		"tickInterval": TickIntervalOneMin,
		"useApi2":      "true",
	}

	c := New("", "")

	fullURI := c.getFullURI(endpoint, params)

	compareURI := "https://bittrex.com/api/v2.0/pub/market/getticks?apikey=&marketName=BTC-LTC&nonce=1513971138&tickInterval=oneMin"

	if fullURI != compareURI {
		t.Errorf("url %s doesn't match expected string %s", fullURI, compareURI)
	}
}

func TestSendRequest(t *testing.T) {
	endpoint := "public/getticker"
	params := map[string]string{
		"market": "BTC-LTC",
	}

	c := New("", "")

	baseResponse := c.sendRequest(endpoint, params)

	if c.err != nil {
		t.Errorf("Base Response %+v\n", baseResponse)
		t.Errorf("Error Message %+v\n", c.err)
	}

}
