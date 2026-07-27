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

type BaseTransactionEvent struct {
	UserId   uuid.UUID
	Currency string
	Amount   string

	Start time.Time
	End   time.Time
	Err   *string
}

type DepositEvent struct {
	BaseTransactionEvent
}

type WithdrawEvent struct {
	BaseTransactionEvent
}

type ExchangeEvent struct {
	BaseTransactionEvent
	ToCurrency       string
	FromExchangeRate string
	ToExchangeRate   string
}

const (
	Deposit  = "Deposit"
	Withdraw = "Withdraw"
	Exchange = "Exchange"
)
