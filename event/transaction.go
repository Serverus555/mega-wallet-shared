package event

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Envelope struct {
	Type  string
	Event json.RawMessage
}

type baseTransactionEvent struct {
	UserId uuid.UUID
	Amount string

	Start time.Time
	End   time.Time
	Err   *string
}

type DepositEvent struct {
	baseTransactionEvent
	Currency string
	Amount   string
}

type WithdrawEvent struct {
	baseTransactionEvent
	Currency string
	Amount   string
}

type ExchangeEvent struct {
	baseTransactionEvent
	FromCurrency     string
	ToCurrency       string
	FromExchangeRate string
	ToExchangeRate   string
}

const (
	Deposit  = "Deposit"
	Withdraw = "Withdraw"
	Exchange = "Exchange"
)
