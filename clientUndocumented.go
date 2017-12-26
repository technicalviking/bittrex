package bittrex

import "encoding/json"

const (
	//TickIntervalOneMin oneMin = 10 days worth of candles
	TickIntervalOneMin = "oneMin"

	//TickIntervalFiveMin fiveMin = 20 days worth of candles
	TickIntervalFiveMin = "fiveMin"

	//TickIntervalThirtyMin thirtyMin = 40 days worth of candles
	TickIntervalThirtyMin = "thirtyMin"

	//TickIntervalHour hour = 60 days worth of candles
	TickIntervalHour = "hour"

	//TickIntervalDay day = 1385 days (nearly four years)
	TickIntervalDay = "day"
)

// PubMarketGetTicks - /pub/market/getticks
// interval must be one of the TickInterval consts
func (c *Client) PubMarketGetTicks(market string, interval string) ([]Candle, error) {
	defer c.clearError()

	params := map[string]string{
		"marketName":   market,
		"tickInterval": interval,
		"useApi2":      "true",
	}

	var parsedResponse *baseResponse

	parsedResponse = c.sendRequest("pub/market/getticks", params)

	if parsedResponse == nil {
		return nil, c.err
	}

	if parsedResponse.Success != true {
		c.setError("api error - pub/market/getticks", parsedResponse.Message)
		return nil, c.err
	}

	var response []Candle

	if err := json.Unmarshal(parsedResponse.Result, &response); err != nil {
		c.setError("parseResponse", err.Error())
		return nil, c.err
	}

	return response, nil
}
