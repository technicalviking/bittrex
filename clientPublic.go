package bittrex

import "encoding/json"

// PublicGetMarkets - public/getmarkets
func (c *Client) PublicGetMarkets() ([]MarketDescription, error) {
	defer c.clearError()

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("public/getmarkets", nil)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - public/getmarkets", parsedResponse.Message)
		return nil, c.err
	}

	var response []MarketDescription

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}

// PublicGetCurrencies - public/getcurrencies
func (c *Client) PublicGetCurrencies() ([]Currency, error) {
	defer c.clearError()

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("public/getcurrencies", nil)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - public/getcurrencies", parsedResponse.Message)
		return nil, c.err
	}

	var response []Currency

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}

// PublicGetTicker - public/getticker
func (c *Client) PublicGetTicker(market string) (Ticker, error) {
	defer c.clearError()

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("public/getticker", map[string]string{"market": market})

	if parsedResponse == nil {
		return Ticker{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - public/getticker", parsedResponse.Message)
		return Ticker{}, c.err
	}

	var response Ticker

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return Ticker{}, c.err
	}

	return response, nil
}

// PublicGetMarketSummaries - public/getmarketsummaries
func (c *Client) PublicGetMarketSummaries() ([]MarketSummary, error) {
	defer c.clearError()

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("public/getmarketsummaries", nil)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - public/getmarketsummaries", parsedResponse.Message)
		return nil, c.err
	}

	var response []MarketSummary

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}

// PublicGetMarketSummary - public/getmarketsummary
func (c *Client) PublicGetMarketSummary(market string) (MarketSummary, error) {
	defer c.clearError()

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("public/getmarketsummary", map[string]string{"market": market})

	if parsedResponse == nil {
		return MarketSummary{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - public/getmarketsummary", parsedResponse.Message)
		return MarketSummary{}, c.err
	}

	var response MarketSummary

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return MarketSummary{}, c.err
	}

	return response, nil
}

// PublicGetOrderBook - public/getorderbook
func (c *Client) PublicGetOrderBook(market string, orderType string) (OrderBook, error) {
	defer c.clearError()

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("public/getorderbook", map[string]string{"market": market, "type": orderType})

	if parsedResponse == nil {
		return OrderBook{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - public/getorderbook", parsedResponse.Message)
		return OrderBook{}, c.err
	}

	var response OrderBook

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return OrderBook{}, c.err
	}

	return response, nil
}

// PublicGetMarketHistory - public/getmarkethistory
func (c *Client) PublicGetMarketHistory(market string) ([]Trade, error) {
	defer c.clearError()

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("public/getmarkethistory", map[string]string{"market": market})

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - public/getmarkethistory", parsedResponse.Message)
		return nil, c.err
	}

	var response []Trade

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}
