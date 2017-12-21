package bittrex

import (
	"fmt"
	"testing"
)

func TestAccountGetBalances(t *testing.T) {
	//If you wanna run these tests, use your own key.
	rClient := New(
		"",
		"",
	)

	var balances []AccountBalance
	var e error

	if balances, e = rClient.AccountGetBalances(); e != nil {
		fmt.Printf("WTF %v\n", e)
		t.Error(e)
	}

	fmt.Printf("BALANCES %v\n", balances)
}
