package accounts

import (
	"github/parking-lot-lld/components"
	"github/parking-lot-lld/components/miscallaneous"
	"github/parking-lot-lld/enums"
)

type ParkingAgent struct {
	id int
	username string
	status enums.AccountStatusEnum
	person miscallaneous.Person
}

func (a ParkingAgent) NewParkingAgent (id int, name string, status enums.AccountStatusEnum, person miscallaneous.Person) *ParkingAgent {
	return &ParkingAgent {
		id: id,
		username: name,
		status: status,
		person: person,
	}
}

func (a ParkingAgent) Login() string {
	return "token"
}

func (a ParkingAgent) Logout() bool {
	return true
}

func (a ParkingAgent) ResetPassword() string {
	return "token"
}

func (a ParkingAgent) ProcessTicket(ticket components.ParkingTicket) {
	// Process Parking Tickets
}