package bittrex

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

// AccountGetBalances - /account/getbalances
func (c *Client) AccountGetBalances() ([]AccountBalance, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey": c.apiKey,
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("account/getbalances", params)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - account/getbalances", parsedResponse.Message)
		return nil, c.err
	}

	var response []AccountBalance

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}

// AccountGetBalance - /account/getbalance
func (c *Client) AccountGetBalance(currency string) (AccountBalance, error) {
	defer c.clearError()

	var parsedResponse *baseResponse

	params := map[string]string{
		"apikey":   c.apiKey,
		"currency": currency,
	}

	parsedResponse = c.sendRequest("account/getbalance", params)

	if parsedResponse == nil {
		return AccountBalance{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - account/getbalance", parsedResponse.Message)
		return AccountBalance{}, c.err
	}

	var response AccountBalance

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return AccountBalance{}, c.err
	}

	return response, nil
}

// AccountGetDepositAddress - /account/getdepositaddress
func (c *Client) AccountGetDepositAddress(currency string) (WalletAddress, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey":   c.apiKey,
		"currency": currency,
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("account/getdepositaddress", params)

	if parsedResponse == nil {
		return WalletAddress{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - account/getdepositaddress", parsedResponse.Message)
		return WalletAddress{}, c.err
	}

	var response WalletAddress

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return WalletAddress{}, c.err
	}

	return response, nil
}

/*
AccountWithdraw - /account/withdraw
paymentId field is optional for the api (used as a memo field for other services
such as CryptoNotes, BitShareX, Nxt).  Set it to empty string to exclude it from
api call
*/
func (c *Client) AccountWithdraw(currency string, quantity decimal.Decimal, address string, paymentID string) (TransactionID, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey":   c.apiKey,
		"currency": currency,
		"quantity": quantity.String(),
		"address":  address,
	}

	if paymentID != "" {
		params["paymentid"] = paymentID
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("account/withdraw", params)

	if parsedResponse == nil {
		return TransactionID{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - account/withdraw", parsedResponse.Message)
		return TransactionID{}, c.err
	}

	var response TransactionID

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return TransactionID{}, c.err
	}

	return response, nil
}

// AccountGetOrder - /account/getorder
func (c *Client) AccountGetOrder(orderID string) (AccountOrderDescription, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey": c.apiKey,
		"uuid":   orderID,
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("account/getorder", params)

	if parsedResponse == nil {
		return AccountOrderDescription{}, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - account/getorder", parsedResponse.Message)
		return AccountOrderDescription{}, c.err
	}

	var response AccountOrderDescription

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return AccountOrderDescription{}, c.err
	}

	return response, nil
}

/*
AccountGetOrderHistory - /account/getorderhistory
market is optional param.  set it to empty strinng to get all markets.
*/
func (c *Client) AccountGetOrderHistory(market string) ([]AccountOrderHistoryDescription, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey": c.apiKey,
	}

	if market != "" {
		params["market"] = market
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("account/getorderhistory", params)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - account/getorderhistory", parsedResponse.Message)
		return nil, c.err
	}

	var response []AccountOrderHistoryDescription

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}

/*
AccountGetWithdrawalHistory - /account/getwithdrawalhistory
setting currency to empty string will get all currencies.
*/
func (c *Client) AccountGetWithdrawalHistory(currency string) ([]TransactionHistoryDescription, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey": c.apiKey,
	}

	if currency != "" {
		params["currency"] = currency
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("account/getwithdrawalhistory", params)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - account/getwithdrawalhistory", parsedResponse.Message)
		return nil, c.err
	}

	var response []TransactionHistoryDescription

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}

/*
AccountGetDepositHistory - /account/getdeposithistory
setting currency to empty string will get all currencies.
*/
func (c *Client) AccountGetDepositHistory(currency string) ([]TransactionHistoryDescription, error) {
	defer c.clearError()

	params := map[string]string{
		"apikey": c.apiKey,
	}

	if currency != "" {
		params["currency"] = currency
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("account/getdeposithistory", params)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - account/getdeposithistory", parsedResponse.Message)
		return nil, c.err
	}

	var response []TransactionHistoryDescription

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}
