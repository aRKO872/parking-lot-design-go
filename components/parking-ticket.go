package components

import (
	"fmt"
	payment "github/parking-lot-lld/components/payment"
	"time"
)

type ParkingTicket struct {
	ticketNo int
	timeOfEntry string
	timeOfExit string
	amount float64
	payment payment.Payment
}

func (pt *ParkingTicket) TicketNo() int {
	return pt.ticketNo
}

func (pt *ParkingTicket) SetTicketNo(ticketNo int) {
	pt.ticketNo = ticketNo
}

func (pt *ParkingTicket) TimeOfEntry() time.Time {
	layout := "2006-01-02 15:04:05.000"
	date, err := time.Parse(layout, pt.timeOfEntry)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Now()
	} else {
		return date
	}
}

func (pt *ParkingTicket) SetTimeOfEntry(timeOfEntry time.Time) {
	dateString := timeOfEntry.Format("2006-01-02 15:04:05.000")
	pt.timeOfEntry = dateString
}

func (pt *ParkingTicket) TimeOfExit() time.Time {
	layout := "2006-01-02 15:04:05.000"
	date, err := time.Parse(layout, pt.timeOfExit)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Now()
	} else {
		return date
	}
}

func (pt *ParkingTicket) SetTimeOfExit(timeOfExit time.Time) {
	dateString := timeOfExit.Format("2006-01-02 15:04:05.000")
	pt.timeOfEntry = dateString
}

func (pt *ParkingTicket) Amount() float64 {
	return pt.amount
}

func (pt *ParkingTicket) SetAmount(amount float64) {
	pt.amount = amount
}

func (pt *ParkingTicket) Payment() payment.Payment {
	return pt.payment
}

func (pt *ParkingTicket) SetPayment(payment payment.Payment) {
	pt.payment = payment
}