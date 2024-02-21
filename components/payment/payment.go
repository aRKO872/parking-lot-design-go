package payment

import (
	types "github/parking-lot-lld/enums"
	"time"
)

type Payment struct {
	id            string
	amount        float64
	status        types.PaymentStatusEnum
	timeOfPayment time.Time
}

func NewPayment(id string, amount float64, status types.PaymentStatusEnum, timeOfPayment time.Time) *Payment {
	return &Payment{
		id:            id,
		amount:        amount,
		status:        status,
		timeOfPayment: timeOfPayment,
	}
}

func (p *Payment) ID() string {
	return p.id
}

func (p *Payment) SetID(id string) {
	p.id = id
}

func (p *Payment) Amount() float64 {
	return p.amount
}

func (p *Payment) SetAmount(amount float64) {
	p.amount = amount
}

func (p *Payment) Status() types.PaymentStatusEnum {
	return p.status
}

func (p *Payment) SetStatus(status types.PaymentStatusEnum) {
	p.status = status
}

func (p *Payment) TimeOfPayment() time.Time {
	return p.timeOfPayment
}

func (p *Payment) SetTimeOfPayment(timeOfPayment time.Time) {
	p.timeOfPayment = timeOfPayment
}