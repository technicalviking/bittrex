package bittrex

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type BittrexTimestamp time.Time

func (bt *BittrexTimestamp) UnmarshalJSON(raw []byte) error {
	var strTimestamp string //"2014-07-09T07:19:30.15"

	if err := json.Unmarshal(raw, &strTimestamp); err != nil {
		return err
	}

	parts := strings.Split(strTimestamp, "T")
	strDate := parts[0]
	strTime := parts[1]

	dateParts := strings.Split(strDate, "-")
	timeParts := strings.Split(strTime, ":")

	var year, month, day, hour, minute, second, nano int

	errs := make([]error, 7)

	year, errs[0] = strconv.Atoi(dateParts[0])
	month, errs[1] = strconv.Atoi(dateParts[1])
	day, errs[2] = strconv.Atoi(dateParts[2])

	hour, errs[3] = strconv.Atoi(timeParts[0])
	minute, errs[4] = strconv.Atoi(timeParts[1])

	secParts := strings.Split(timeParts[2], ".")

	second, errs[5] = strconv.Atoi(secParts[0])

	if len(secParts) > 1 {
		nano, errs[6] = strconv.Atoi(secParts[1])
		nano *= int(math.Pow10(8 - (len(secParts[1]) - 1)))
	} else {
		nano = 0
	}

	newTime := time.Date(
		year,
		time.Month(month),
		day,
		hour,
		minute,
		second,
		nano,
		time.UTC,
	)

	*bt = BittrexTimestamp(newTime)

	return nil
}

func (bt *BittrexTimestamp) String() string {
	cast := time.Time(*bt)
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", cast.Year(), cast.Month(), cast.Day(), cast.Hour(), cast.Minute(), cast.Second())
}

//MarketDescription Result element as described under /public/getmarkets
type MarketDescription struct {
	MarketCurrency     string           `json:"MarketCurrency"`
	BaseCurrency       string           `json:"BaseCurrency"`
	MarketCurrencyLong string           `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string           `json:"BaseCurrencyLong"`
	MinTradeSize       *big.Float       `json:"MinTradeSize"`
	MarketName         string           `json:"MarketName"`
	IsActive           bool             `json:"IsActive"`
	Created            BittrexTimestamp `json:"Created"`
}

//Currency Result element as described under /public/getcurrencies
type Currency struct {
	Currency        string     `json:"Currency"`
	CurrencyLong    string     `json:"CurrencyLong"`
	MinConfirmation int        `json:"MinConfirmation"`
	TxFee           *big.Float `json:"TxFee"`
	IsActive        bool       `json:"IsActive"`
	CoinType        string     `json:"CoinType"`
	BaseAddress     string     `json:"BaseAddress"`
}

//Ticker Result element as described under /public/getticker
type Ticker struct {
	Bid  *big.Float `json:"Bid"`
	Ask  *big.Float `json:"Ask"`
	Last *big.Float `json:"Last"`
}

//MarketSummary result element as described under /public/getmarketsummaries
type MarketSummary struct {
	MarketName        string           `json:"MarketName"`        // : "BTC-888",
	High              *big.Float       `json:"High"`              // : 0.00000919,
	Low               *big.Float       `json:"Low"`               // : 0.00000820,
	Volume            *big.Float       `json:"Volume"`            // : 74339.61396015,
	Last              *big.Float       `json:"Last"`              // : 0.00000820,
	BaseVolume        *big.Float       `json:"BaseVolume"`        // : 0.64966963,
	TimeStamp         BittrexTimestamp `json:"TimeStamp"`         // : "2014-07-09T07:19:30.15",
	Bid               *big.Float       `json:"Bid"`               // : 0.00000820,
	Ask               *big.Float       `json:"Ask"`               // : 0.00000831,
	OpenBuyOrders     int              `json:"OpenBuyOrders"`     // : 15,
	OpenSellOrders    int              `json:"OpenSellOrders"`    // : 15,
	PrevDay           *big.Float       `json:"PrevDay"`           // : 0.00000821,
	Created           BittrexTimestamp `json:"Created"`           // : "2014-03-20T06:00:00",
	DisplayMarketName string           `json:"DisplayMarketName"` // : null
}

//OrderElement element found under 'buy' or 'sell' in an OrderBook
type OrderElement struct {
	Quantity *big.Float `json:"Quantity"`
	Rate     *big.Float `json:"Rate"`
}

//OrderBook Result body of /public/getorderbook
type OrderBook struct {
	Buy  []OrderElement `json:"buy"`
	Sell []OrderElement `json:"sell"`
}

//Trade result element as described under /public/getmarkethistory
type Trade struct {
	ID        string           `json:"Id"`        // : 319435,
	TimeStamp BittrexTimestamp `json:"TimeStamp"` // : "2014-07-09T03:21:20.08",
	Quantity  *big.Float       `json:"Quantity"`  // : 0.30802438,
	Price     *big.Float       `json:"Price"`     // : 0.01263400,
	Total     *big.Float       `json:"Total"`     // : 0.00389158,
	FillType  string           `json:"FillType"`  // : "FILL",
	OrderType string           `json:"OrderType"` // : "BUY" or "SELL"
}

//TransactionID Result body of /market/buylimit and /market/sellimit
type TransactionID struct {
	UUID string `json:"uuid"`
}

//OrderDescription result element as described under /market/getopenorders
type OrderDescription struct {
	UUID              string           `json:"Uuid"`              // : null,
	OrderUUID         string           `json:"OrderUuid"`         // : "09aa5bb6-8232-41aa-9b78-a5a1093e0211",
	Exchange          string           `json:"Exchange"`          // : "BTC-LTC",
	OrderType         string           `json:"OrderType"`         // : "LIMIT_SELL",
	Quantity          *big.Float       `json:"Quantity"`          // : 5.00000000,
	QuantityRemaining *big.Float       `json:"QuantityRemaining"` // : 5.00000000,
	Limit             *big.Float       `json:"Limit"`             // : 2.00000000,
	CommissionPaid    *big.Float       `json:"CommissionPaid"`    // : 0.00000000,
	Price             *big.Float       `json:"Price"`             // : 0.00000000,
	PricePerUnit      *big.Float       `json:"PricePerUnit"`      // : null,
	Opened            BittrexTimestamp `json:"Opened"`            // : "2014-07-09T03:55:48.77",
	Closed            BittrexTimestamp `json:"Closed"`            // : null,
	CancelInitiated   bool             `json:"CancelInitiated"`   // : false,
	ImmediateOrCancel bool             `json:"ImmediateOrCancel"` // : false,
	IsConditional     bool             `json:"IsConditional"`     // : false,
	Condition         string           `json:"Condition"`         // : null,
	ConditionTarget   string           `json:"ConditionTarget"`   // : null
}

//AccountOrderDescription result body of /account/getorder
type AccountOrderDescription struct {
	AccountID                  string           `json:"AccountId"`                  // : null,
	OrderUUID                  string           `json:"OrderUuid"`                  // : "0cb4c4e4-bdc7-4e13-8c13-430e587d2cc1",
	Exchange                   string           `json:"Exchange"`                   // : "BTC-SHLD",
	Type                       string           `json:"Type"`                       // : "LIMIT_BUY",
	Quantity                   *big.Float       `json:"Quantity"`                   // : 1000.00000000,
	QuantityRemaining          *big.Float       `json:"QuantityRemaining"`          // : 1000.00000000,
	Limit                      *big.Float       `json:"Limit"`                      // : 0.00000001,
	Reserved                   *big.Float       `json:"Reserved"`                   // : 0.00001000,
	ReserveRemaining           *big.Float       `json:"ReserveRemaining"`           // : 0.00001000,
	CommissionReserved         *big.Float       `json:"CommissionReserved"`         // : 0.00000002,
	CommissionReserveRemaining *big.Float       `json:"CommissionReserveRemaining"` // : 0.00000002,
	CommissionPaid             *big.Float       `json:"CommissionPaid"`             // : 0.00000000,
	Price                      *big.Float       `json:"Price"`                      // : 0.00000000,
	PricePerUnit               *big.Float       `json:"PricePerUnit"`               // : null,
	Opened                     BittrexTimestamp `json:"Opened"`                     // : "2014-07-13T07:45:46.27",
	Closed                     BittrexTimestamp `json:"Closed"`                     // : null,
	IsOpen                     bool             `json:"IsOpen"`                     // : true,
	Sentinel                   string           `json:"Sentinel"`                   // : "6c454604-22e2-4fb4-892e-179eede20972",
	CancelInitiated            bool             `json:"CancelInitiated"`            // : false,
	ImmediateOrCancel          bool             `json:"ImmediateOrCancel"`          // : false,
	IsConditional              bool             `json:"IsConditional"`              // : false,
	Condition                  string           `json:"Condition"`                  // : "NONE",
	ConditionTarget            string           `json:"ConditionTarget"`            // : null
}

//AccountOrderHistoryDescription result element of /account/getorderhistory
type AccountOrderHistoryDescription struct {
	OrderUUID         string           `json:"OrderUuid"`         // : "fd97d393-e9b9-4dd1-9dbf-f288fc72a185",
	Exchange          string           `json:"Exchange"`          // : "BTC-LTC",
	TimeStamp         BittrexTimestamp `json:"TimeStamp"`         // : "2014-07-09T04:01:00.667",
	OrderType         string           `json:"OrderType"`         // : "LIMIT_BUY",
	Limit             *big.Float       `json:"Limit"`             // : 0.00000001,
	Quantity          *big.Float       `json:"Quantity"`          // : 100000.00000000,
	QuantityRemaining *big.Float       `json:"QuantityRemaining"` // : 100000.00000000,
	Commission        *big.Float       `json:"Commission"`        // : 0.00000000,
	Price             *big.Float       `json:"Price"`             // : 0.00000000,
	PricePerUnit      *big.Float       `json:"PricePerUnit"`      // : null,
	IsConditional     bool             `json:"IsConditional"`     // : false,
	Condition         string           `json:"Condition"`         // : null,
	ConditionTarget   string           `json:"ConditionTarget"`   // : null,
	ImmediateOrCancel bool             `json:"ImmediateOrCancel"` // : false
}

//TransactionHistoryDescription result element of /account/getwithdrawalhistory and /account/getdeposithistory
type TransactionHistoryDescription struct {
	PaymentUUID    string           `json:"PaymentUuid"`    // : "b52c7a5c-90c6-4c6e-835c-e16df12708b1",
	Currency       string           `json:"Currency"`       // : "BTC",
	Amount         *big.Float       `json:"Amount"`         // : 17.00000000,
	Address        string           `json:"Address"`        // : "1DeaaFBdbB5nrHj87x3NHS4onvw1GPNyAu",
	Opened         BittrexTimestamp `json:"Opened"`         // : "2014-07-09T04:24:47.217",
	Authorized     bool             `json:"Authorized"`     // : true,
	PendingPayment bool             `json:"PendingPayment"` // : false,
	TxCost         *big.Float       `json:"TxCost"`         // : 0.00020000,
	TxID           string           `json:"TxId"`           // : null,
	Canceled       bool             `json:"Canceled"`       // : true,
	InvalidAddress bool             `json:"InvalidAddress"` // : false
}

//AccountBalance result element as described under /account/getbalances. also the result body of /account/getbalance
type AccountBalance struct {
	Currency      string     `json:"Currency"`      // : "DOGE",
	Balance       *big.Float `json:"Balance"`       // : 0.00000000,
	Available     *big.Float `json:"Available"`     // : 0.00000000,
	Pending       *big.Float `json:"Pending"`       // : 0.00000000,
	CryptoAddress string     `json:"CryptoAddress"` // : "DLxcEt3AatMyr2NTatzjsfHNoB9NT62HiF",
	Requested     bool       `json:"Requested"`     // : false,
	UUID          string     `json:"Uuid"`          // : null
}

//WalletAddress result body of /account/getdepositaddress
type WalletAddress struct {
	Currency string `json:"Currency"` // : "VTC"
	Address  string `json:"Address"`  // : "Vy5SKeKGXUHKS2WVpJ76HYuKAu3URastUo"
}

//Candle result element as described under v2.0/pub/market/getticks
type Candle struct {
	TimeStamp BittrexTimestamp `json:"T"`
	Open      *big.Float       `json:"O"`
	Close     *big.Float       `json:"C"`
	High      *big.Float       `json:"H"`
	Low       *big.Float       `json:"L"`
	//Volume amount traded in the altcoin (Ex: the LTC in BTC-LTC)
	Volume *big.Float `json:"V"`
	//Volume amount traded in the base coin (Ex: the BTC in BTC-LTC)
	BaseVolume *big.Float `json:"BV"`
}

//OrderUpdate Update to an order listed under buys and sells in ExchangeState
type OrderUpdate struct {
	OrderElement //embed
	Type         int
}

//Fill structure found inside an ExchangeState object
type Fill struct {
	OrderElement //embed
	OrderType    string
	Timestamp    BittrexTimestamp
}

// ExchangeState contains fills and order book updates for a market.
type ExchangeState struct {
	MarketName string
	Nounce     int
	Buys       []OrderUpdate
	Sells      []OrderUpdate
	Fills      []Fill
	Initial    bool
}
