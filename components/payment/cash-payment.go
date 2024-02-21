package payment

import (
	types "github/parking-lot-lld/enums"
	"time"
)

// Cash Struct Inherited from PaymentByCash 
type PaymentByCash struct {
	Payment
}

func NewPaymentByCash(id string, amount float64, status types.PaymentStatusEnum, timeOfPayment time.Time) *PaymentByCash {
	return &PaymentByCash{
		Payment: *NewPayment(id, amount, status, timeOfPayment),
	}
}

func (p *PaymentByCash) InitiateTransaction() {
	// Payment by Cash Logic
}