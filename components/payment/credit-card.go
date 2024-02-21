package payment

import (
	types "github/parking-lot-lld/enums"
	"time"
)

// Credit Card Struct Inherited from PaymentByCash
type PaymentByCreditCard struct {
	Payment
}

func NewPaymentByCreditCard(id string, amount float64, status types.PaymentStatusEnum, timeOfPayment time.Time) *PaymentByCreditCard {
	return &PaymentByCreditCard{
		Payment: *NewPayment(id, amount, status, timeOfPayment),
	}
}

func (p *PaymentByCreditCard) InitiateTransaction() {
	// Payment by Credit Card Logic
}