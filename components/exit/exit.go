package exit

import "github/parking-lot-lld/components"

type Exit struct {
	id int
}

func (e Exit) NewExit(id int) *Exit {
	return &Exit{
		id: id,
	}
}

func (e Exit) Id() int {
	return e.id
} 

func (e *Exit) SetId(id int) {
	e.id = id
}

func (e Exit) ScanAndPayTicket(*components.ParkingTicket) {
	// Add Logic to Create and Return a Parking Ticket
}