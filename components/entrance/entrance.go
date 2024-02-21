package entrance

import "github/parking-lot-lld/components"

type Entrance struct {
	id int
}

func (e Entrance) NewEntrance(id int) *Entrance {
	return &Entrance{
		id: id,
	}
}

func (e Entrance) Id() int {
	return e.id
} 

func (e *Entrance) SetId(id int) {
	e.id = id
}

func (e Entrance) GetTicket() *components.ParkingTicket {
	// Add Logic to Create and Return a Parking Ticket
	return nil
}