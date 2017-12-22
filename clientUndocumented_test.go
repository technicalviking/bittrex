package bittrex

import (
	"fmt"
	"testing"
)

func TestPubMarketGetTicks(t *testing.T) {
	//If you wanna run these tests, use your own key.
	rClient := New(
		"",
		"",
	)

	var candles []Candle
	var e error

	if candles, e = rClient.PubMarketGetTicks("BTC-LTC", TickIntervalOneMin); e != nil {
		fmt.Printf("WTF %v\n", e)
		t.Error(e)
	}

	fmt.Printf("candles %v\n", candles)
}
