package bittrex

import "encoding/json"

const (
	//TickIntervalOneMin oneMin
	TickIntervalOneMin = "oneMin"

	//TickIntervalFiveMin fiveMin
	TickIntervalFiveMin = "fiveMin"

	//TickIntervalThirtyMin thirtyMin
	TickIntervalThirtyMin = "thirtyMin"

	//TickIntervalHour hour
	TickIntervalHour = "hour"

	//TickIntervalDay day
	TickIntervalDay = "day"
)

// PubMarketGetTicks - /pub/market/getticks
// interval must be one of the TickInterval consts
func (c *Client) PubMarketGetTicks(market string, interval string) ([]Candle, error) {
	defer c.clearError()

	params := map[string]string{
		"market":   market,
		"interval": interval,
		"useApi2":  "true",
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("/pub/market/getticks", params)

	if parsedResponse.Success != true {
		c.setError("api error - /pub/market/getticks", parsedResponse.Message)
		return nil, c.err
	}

	var response []Candle

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}
