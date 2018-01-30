package bittrex

import (
	"encoding/json"
	"math/big"
)

// MarketBuyLimit - market/buylimit
func (c *Client) MarketBuyLimit(market string, quantity *big.Float, rate *big.Float) (TransactionID, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey":   c.apiKey,
		"market":   market,
		"quantity": quantity.String(),
		"rate":     rate.String(),
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("market/buylimit", params)

	if parsedResponse == nil {
		return TransactionID{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - market/buylimit", parsedResponse.Message)
		return TransactionID{}, c.err
	}

	var response TransactionID

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return TransactionID{}, c.err
	}

	return response, nil
}

// MarketSellLimit - market/selllimit
func (c *Client) MarketSellLimit(market string, quantity *big.Float, rate *big.Float) (TransactionID, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey":   c.apiKey,
		"market":   market,
		"quantity": quantity.String(),
		"rate":     rate.String(),
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("market/selllimit", params)

	if parsedResponse == nil {
		return TransactionID{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - market/selllimit", parsedResponse.Message)
		return TransactionID{}, c.err
	}

	var response TransactionID

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return TransactionID{}, c.err
	}

	return response, nil
}

// MarketBuyMarket - market/buymarket - EXPERIMENTAL
func (c *Client) MarketBuyMarket(market string, quantity *big.Float, rate *big.Float) (TransactionID, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey":   c.apiKey,
		"market":   market,
		"quantity": quantity.String(),
		"rate":     rate.String(),
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("market/buymarket", params)

	if parsedResponse == nil {
		return TransactionID{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - market/buymarket", parsedResponse.Message)
		return TransactionID{}, c.err
	}

	var response TransactionID

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return TransactionID{}, c.err
	}

	return response, nil
}

// MarketSellMarket - market/sellmarket - EXPERIMENTAL
func (c *Client) MarketSellMarket(market string, quantity *big.Float, rate *big.Float) (TransactionID, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey":   c.apiKey,
		"market":   market,
		"quantity": quantity.String(),
		"rate":     rate.String(),
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("market/sellmarket", params)

	if parsedResponse == nil {
		return TransactionID{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - market/sellmarket", parsedResponse.Message)
		return TransactionID{}, c.err
	}

	var response TransactionID

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return TransactionID{}, c.err
	}

	return response, nil
}

// MarketCancel - market/cancel
func (c *Client) MarketCancel(uuid string) (bool, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey": c.apiKey,
		"uuid":   uuid,
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("market/cancel", params)

	if parsedResponse == nil {
		return false, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - market/cancel", parsedResponse.Message)
		return false, c.err
	}

	return true, nil
}

// MarketGetOpenOrders - market/getopenorders
func (c *Client) MarketGetOpenOrders(market string) ([]OrderDescription, error) {
	defer c.clearError()

	params := map[string]string{
		"market": market,
		"apikey": c.apiKey,
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("market/getopenorders", params)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - market/getopenorders", parsedResponse.Message)
		return nil, c.err
	}

	var response []OrderDescription

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}
