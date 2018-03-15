package bittrex

import (
	"fmt"
	"testing"
)

var (
	//If you wanna run these tests, use your own key.
	rClient = New(
		"",
		"",
	)
)

func TestAccountGetBalances(t *testing.T) {

	var balances []AccountBalance
	var e error

	if balances, e = rClient.AccountGetBalances(); e != nil {
		fmt.Printf("WTF %v\n", e)
		t.Error(e)
	}

	t.Errorf("BALANCES %v\n", balances)
}

func TestAccountGetBalance(t *testing.T) {

	var balance AccountBalance
	var e error

	if balance, e = rClient.AccountGetBalance("NBT"); e != nil {
		fmt.Printf("WTF %v\n", e)
		t.Error(e)
	}

	t.Errorf("BALANCE %v\n", balance)
}
