package bittrex

import (
	"fmt"
	"testing"
)

func TestWsSubExchangeUpdates(t *testing.T) {
	bittrexClient := New("", "")

	sub := bittrexClient.WsSubExchangeUpdates("")

	for i := 0; i < 5; i++ {
		select {
		case d := <-sub.Data:
			fmt.Printf("What is this: %v\n", d)
		case e := <-sub.Error:
			t.Errorf("Error lol %s \n", e.Error())
			i = 999
		}
	}

	close(sub.Done)
}
