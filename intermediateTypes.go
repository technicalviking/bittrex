package bittrex

import (
	"encoding/json"
	"math/big"
)

func castToBigFloat(num json.Number) (result *big.Float) {
	result, _, _ = big.ParseFloat(num.String(), 10, uint(len(num)), big.ToNearestEven)
	return
}

func (m *MarketDescription) UnmarshalJSON(raw []byte) error {
	temp := struct {
		MarketCurrency     string           `json:"MarketCurrency"`
		BaseCurrency       string           `json:"BaseCurrency"`
		MarketCurrencyLong string           `json:"MarketCurrencyLong"`
		BaseCurrencyLong   string           `json:"BaseCurrencyLong"`
		MinTradeSize       json.Number      `json:"MinTradeSize"`
		MarketName         string           `json:"MarketName"`
		IsActive           bool             `json:"IsActive"`
		Created            BittrexTimestamp `json:"Created"`
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = MarketDescription{
		temp.MarketCurrency,
		temp.BaseCurrency,
		temp.MarketCurrencyLong,
		temp.BaseCurrencyLong,
		castToBigFloat(temp.MinTradeSize),
		temp.MarketName,
		temp.IsActive,
		temp.Created,
	}

	return nil
}

func (m *Currency) UnmarshalJSON(raw []byte) error {
	temp := struct {
		Currency        string      `json:"Currency"`
		CurrencyLong    string      `json:"CurrencyLong"`
		MinConfirmation int         `json:"MinConfirmation"`
		TxFee           json.Number `json:"TxFee"`
		IsActive        bool        `json:"IsActive"`
		CoinType        string      `json:"CoinType"`
		BaseAddress     string      `json:"BaseAddress"`
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = Currency{
		Currency:        temp.Currency,
		CurrencyLong:    temp.CurrencyLong,
		MinConfirmation: temp.MinConfirmation,
		TxFee:           castToBigFloat(temp.TxFee),
		IsActive:        temp.IsActive,
		CoinType:        temp.CoinType,
		BaseAddress:     temp.BaseAddress,
	}

	return nil
}

func (m *Ticker) UnmarshalJSON(raw []byte) error {
	temp := struct {
		Bid  json.Number `json:"Bid"`
		Ask  json.Number `json:"Ask"`
		Last json.Number `json:"Last"`
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = Ticker{
		Bid:  castToBigFloat(temp.Bid),
		Ask:  castToBigFloat(temp.Ask),
		Last: castToBigFloat(temp.Last),
	}

	return nil
}

func (m *MarketSummary) UnmarshalJSON(raw []byte) error {
	temp := struct {
		MarketName        string           `json:"MarketName"`          // : "BTC-888",
		High              json.Number      `json:"High"`                // : 0.00000919,
		Low               json.Number      `json:"Low"`                 // : 0.00000820,
		Volume            json.Number      `json:"Volume"`              // : 74339.61396015,
		Last              json.Number      `json:"Last"`                // : 0.00000820,
		BaseVolume        json.Number      `json:"BaseVolume"`          // : 0.64966963,
		TimeStamp         BittrexTimestamp `json:"TimeStamp,omitempty"` // : "2014-07-09T07:19:30.15",
		Bid               json.Number      `json:"Bid"`                 // : 0.00000820,
		Ask               json.Number      `json:"Ask"`                 // : 0.00000831,
		OpenBuyOrders     int              `json:"OpenBuyOrders"`       // : 15,
		OpenSellOrders    int              `json:"OpenSellOrders"`      // : 15,
		PrevDay           json.Number      `json:"PrevDay"`             // : 0.00000821,
		Created           BittrexTimestamp `json:"Created,omitempty"`   // : "2014-03-20T06:00:00",
		DisplayMarketName string           `json:"DisplayMarketName"`   // : null
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = MarketSummary{
		MarketName:        temp.MarketName,
		High:              castToBigFloat(temp.High),
		Low:               castToBigFloat(temp.Low),
		Volume:            castToBigFloat(temp.Volume),
		Last:              castToBigFloat(temp.Last),
		BaseVolume:        castToBigFloat(temp.BaseVolume),
		TimeStamp:         temp.TimeStamp,
		Bid:               castToBigFloat(temp.Bid),
		Ask:               castToBigFloat(temp.Ask),
		OpenBuyOrders:     temp.OpenBuyOrders,
		OpenSellOrders:    temp.OpenSellOrders,
		PrevDay:           castToBigFloat(temp.PrevDay),
		Created:           temp.Created,
		DisplayMarketName: temp.DisplayMarketName,
	}

	return nil
}

func (m *OrderElement) UnmarshalJSON(raw []byte) error {
	temp := struct {
		Quantity json.Number `json:"Quantity"`
		Rate     json.Number `json:"Rate"`
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = OrderElement{
		Quantity: castToBigFloat(temp.Quantity),
		Rate:     castToBigFloat(temp.Rate),
	}

	return nil
}

func (m *Trade) UnmarshalJSON(raw []byte) error {
	temp := struct {
		ID        string           `json:"Id"`        // : 319435,
		TimeStamp BittrexTimestamp `json:"TimeStamp"` // : "2014-07-09T03:21:20.08",
		Quantity  json.Number      `json:"Quantity"`  // : 0.30802438,
		Price     json.Number      `json:"Price"`     // : 0.01263400,
		Total     json.Number      `json:"Total"`     // : 0.00389158,
		FillType  string           `json:"FillType"`  // : "FILL",
		OrderType string           `json:"OrderType"` // : "BUY" or "SELL"
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = Trade{
		ID:        temp.ID,
		TimeStamp: temp.TimeStamp,
		Quantity:  castToBigFloat(temp.Quantity),
		Price:     castToBigFloat(temp.Price),
		Total:     castToBigFloat(temp.Total),
		FillType:  temp.FillType,
		OrderType: temp.OrderType,
	}

	return nil
}

func (m *OrderDescription) UnmarshalJSON(raw []byte) error {
	temp := struct {
		UUID              string           `json:"Uuid"`              // : null,
		OrderUUID         string           `json:"OrderUuid"`         // : "09aa5bb6-8232-41aa-9b78-a5a1093e0211",
		Exchange          string           `json:"Exchange"`          // : "BTC-LTC",
		OrderType         string           `json:"OrderType"`         // : "LIMIT_SELL",
		Quantity          json.Number      `json:"Quantity"`          // : 5.00000000,
		QuantityRemaining json.Number      `json:"QuantityRemaining"` // : 5.00000000,
		Limit             json.Number      `json:"Limit"`             // : 2.00000000,
		CommissionPaid    json.Number      `json:"CommissionPaid"`    // : 0.00000000,
		Price             json.Number      `json:"Price"`             // : 0.00000000,
		PricePerUnit      json.Number      `json:"PricePerUnit"`      // : null,
		Opened            BittrexTimestamp `json:"Opened"`            // : "2014-07-09T03:55:48.77",
		Closed            BittrexTimestamp `json:"Closed,omitempty"`  // : null,
		CancelInitiated   bool             `json:"CancelInitiated"`   // : false,
		ImmediateOrCancel bool             `json:"ImmediateOrCancel"` // : false,
		IsConditional     bool             `json:"IsConditional"`     // : false,
		Condition         string           `json:"Condition"`         // : null,
		ConditionTarget   string           `json:"ConditionTarget"`   // : null
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = OrderDescription{
		UUID:              temp.UUID,
		OrderUUID:         temp.OrderUUID,
		Exchange:          temp.Exchange,
		OrderType:         temp.OrderType,
		Quantity:          castToBigFloat(temp.Quantity),
		QuantityRemaining: castToBigFloat(temp.QuantityRemaining),
		Limit:             castToBigFloat(temp.Limit),
		CommissionPaid:    castToBigFloat(temp.CommissionPaid),
		Price:             castToBigFloat(temp.Price),
		PricePerUnit:      castToBigFloat(temp.PricePerUnit),
		Opened:            temp.Opened,
		Closed:            temp.Closed,
		CancelInitiated:   temp.CancelInitiated,
		ImmediateOrCancel: temp.ImmediateOrCancel,
		IsConditional:     temp.IsConditional,
		Condition:         temp.Condition,
		ConditionTarget:   temp.ConditionTarget,
	}

	return nil
}

func (m *AccountOrderDescription) UnmarshalJSON(raw []byte) error {
	temp := struct {
		AccountID                  string           `json:"AccountId"`                  // : null,
		OrderUUID                  string           `json:"OrderUuid"`                  // : "0cb4c4e4-bdc7-4e13-8c13-430e587d2cc1",
		Exchange                   string           `json:"Exchange"`                   // : "BTC-SHLD",
		Type                       string           `json:"Type"`                       // : "LIMIT_BUY",
		Quantity                   json.Number      `json:"Quantity"`                   // : 1000.00000000,
		QuantityRemaining          json.Number      `json:"QuantityRemaining"`          // : 1000.00000000,
		Limit                      json.Number      `json:"Limit"`                      // : 0.00000001,
		Reserved                   json.Number      `json:"Reserved"`                   // : 0.00001000,
		ReserveRemaining           json.Number      `json:"ReserveRemaining"`           // : 0.00001000,
		CommissionReserved         json.Number      `json:"CommissionReserved"`         // : 0.00000002,
		CommissionReserveRemaining json.Number      `json:"CommissionReserveRemaining"` // : 0.00000002,
		CommissionPaid             json.Number      `json:"CommissionPaid"`             // : 0.00000000,
		Price                      json.Number      `json:"Price"`                      // : 0.00000000,
		PricePerUnit               json.Number      `json:"PricePerUnit"`               // : null,
		Opened                     BittrexTimestamp `json:"Opened"`                     // : "2014-07-13T07:45:46.27",
		Closed                     BittrexTimestamp `json:"Closed,omitempty"`           // : null,
		IsOpen                     bool             `json:"IsOpen"`                     // : true,
		Sentinel                   string           `json:"Sentinel"`                   // : "6c454604-22e2-4fb4-892e-179eede20972",
		CancelInitiated            bool             `json:"CancelInitiated"`            // : false,
		ImmediateOrCancel          bool             `json:"ImmediateOrCancel"`          // : false,
		IsConditional              bool             `json:"IsConditional"`              // : false,
		Condition                  string           `json:"Condition"`                  // : "NONE",
		ConditionTarget            string           `json:"ConditionTarget"`            // : null
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = AccountOrderDescription{
		AccountID:                  temp.AccountID,
		OrderUUID:                  temp.OrderUUID,
		Exchange:                   temp.Exchange,
		Type:                       temp.Type,
		Quantity:                   castToBigFloat(temp.Quantity),
		QuantityRemaining:          castToBigFloat(temp.QuantityRemaining),
		Limit:                      castToBigFloat(temp.Limit),
		Reserved:                   castToBigFloat(temp.Reserved),
		ReserveRemaining:           castToBigFloat(temp.ReserveRemaining),
		CommissionReserved:         castToBigFloat(temp.CommissionReserved),
		CommissionReserveRemaining: castToBigFloat(temp.CommissionReserveRemaining),
		CommissionPaid:             castToBigFloat(temp.CommissionPaid),
		Price:                      castToBigFloat(temp.Price),
		PricePerUnit:               castToBigFloat(temp.PricePerUnit),
		Opened:                     temp.Opened,
		Closed:                     temp.Closed,
		IsOpen:                     temp.IsOpen,
		Sentinel:                   temp.Sentinel,
		CancelInitiated:            temp.CancelInitiated,
		ImmediateOrCancel:          temp.ImmediateOrCancel,
		IsConditional:              temp.IsConditional,
		Condition:                  temp.Condition,
		ConditionTarget:            temp.ConditionTarget,
	}

	return nil
}

func (m *AccountOrderHistoryDescription) UnmarshalJSON(raw []byte) error {
	temp := struct {
		OrderUUID         string           `json:"OrderUuid"`         // : "fd97d393-e9b9-4dd1-9dbf-f288fc72a185",
		Exchange          string           `json:"Exchange"`          // : "BTC-LTC",
		TimeStamp         BittrexTimestamp `json:"TimeStamp"`         // : "2014-07-09T04:01:00.667",
		OrderType         string           `json:"OrderType"`         // : "LIMIT_BUY",
		Limit             json.Number      `json:"Limit"`             // : 0.00000001,
		Quantity          json.Number      `json:"Quantity"`          // : 100000.00000000,
		QuantityRemaining json.Number      `json:"QuantityRemaining"` // : 100000.00000000,
		Commission        json.Number      `json:"Commission"`        // : 0.00000000,
		Price             json.Number      `json:"Price"`             // : 0.00000000,
		PricePerUnit      json.Number      `json:"PricePerUnit"`      // : null,
		IsConditional     bool             `json:"IsConditional"`     // : false,
		Condition         string           `json:"Condition"`         // : null,
		ConditionTarget   string           `json:"ConditionTarget"`   // : null,
		ImmediateOrCancel bool             `json:"ImmediateOrCancel"` // : false
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = AccountOrderHistoryDescription{
		OrderUUID:         temp.OrderUUID,
		Exchange:          temp.Exchange,
		TimeStamp:         temp.TimeStamp,
		OrderType:         temp.OrderType,
		Limit:             castToBigFloat(temp.Limit),
		Quantity:          castToBigFloat(temp.Quantity),
		QuantityRemaining: castToBigFloat(temp.QuantityRemaining),
		Commission:        castToBigFloat(temp.Commission),
		Price:             castToBigFloat(temp.Price),
		PricePerUnit:      castToBigFloat(temp.PricePerUnit),
		IsConditional:     temp.IsConditional,
		Condition:         temp.Condition,
		ConditionTarget:   temp.ConditionTarget,
		ImmediateOrCancel: temp.ImmediateOrCancel,
	}

	return nil
}

func (m *TransactionHistoryDescription) UnmarshalJSON(raw []byte) error {
	temp := struct {
		PaymentUUID    string           `json:"PaymentUuid"`    // : "b52c7a5c-90c6-4c6e-835c-e16df12708b1",
		Currency       string           `json:"Currency"`       // : "BTC",
		Amount         json.Number      `json:"Amount"`         // : 17.00000000,
		Address        string           `json:"Address"`        // : "1DeaaFBdbB5nrHj87x3NHS4onvw1GPNyAu",
		Opened         BittrexTimestamp `json:"Opened"`         // : "2014-07-09T04:24:47.217",
		Authorized     bool             `json:"Authorized"`     // : true,
		PendingPayment bool             `json:"PendingPayment"` // : false,
		TxCost         json.Number      `json:"TxCost"`         // : 0.00020000,
		TxID           string           `json:"TxId"`           // : null,
		Canceled       bool             `json:"Canceled"`       // : true,
		InvalidAddress bool             `json:"InvalidAddress"` // : false
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = TransactionHistoryDescription{
		PaymentUUID:    temp.PaymentUUID,
		Currency:       temp.Currency,
		Amount:         castToBigFloat(temp.Amount),
		Address:        temp.Address,
		Opened:         temp.Opened,
		Authorized:     temp.Authorized,
		PendingPayment: temp.PendingPayment,
		TxCost:         castToBigFloat(temp.TxCost),
		TxID:           temp.TxID,
		Canceled:       temp.Canceled,
		InvalidAddress: temp.InvalidAddress,
	}

	return nil
}

func (m *AccountBalance) UnmarshalJSON(raw []byte) error {
	temp := struct {
		Currency      string      `json:"Currency"`      // : "DOGE",
		Balance       json.Number `json:"Balance"`       // : 0.00000000,
		Available     json.Number `json:"Available"`     // : 0.00000000,
		Pending       json.Number `json:"Pending"`       // : 0.00000000,
		CryptoAddress string      `json:"CryptoAddress"` // : "DLxcEt3AatMyr2NTatzjsfHNoB9NT62HiF",
		Requested     bool        `json:"Requested"`     // : false,
		UUID          string      `json:"Uuid"`          // : null
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = AccountBalance{
		Currency:      temp.Currency,
		Balance:       castToBigFloat(temp.Balance),
		Available:     castToBigFloat(temp.Available),
		Pending:       castToBigFloat(temp.Pending),
		CryptoAddress: temp.CryptoAddress,
		Requested:     temp.Requested,
		UUID:          temp.UUID,
	}

	return nil
}

func (m *Candle) UnmarshalJSON(raw []byte) error {
	temp := struct {
		TimeStamp BittrexTimestamp `json:"T"`
		Open      json.Number      `json:"O"`
		Close     json.Number      `json:"C"`
		High      json.Number      `json:"H"`
		Low       json.Number      `json:"L"`
		//Volume amount traded in the altcoin (Ex: the LTC in BTC-LTC)
		Volume json.Number `json:"V"`
		//Volume amount traded in the base coin (Ex: the BTC in BTC-LTC)
		BaseVolume json.Number `json:"BV"`
	}{}

	if err := json.Unmarshal(raw, &temp); err != nil {
		return err
	}

	*m = Candle{
		TimeStamp:  temp.TimeStamp,
		Open:       castToBigFloat(temp.Open),
		Close:      castToBigFloat(temp.Close),
		High:       castToBigFloat(temp.High),
		Low:        castToBigFloat(temp.Low),
		Volume:     castToBigFloat(temp.Volume),
		BaseVolume: castToBigFloat(temp.BaseVolume),
	}

	return nil
}
