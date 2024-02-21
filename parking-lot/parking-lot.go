package parking_lot

import (
	"github/parking-lot-lld/components"
	"github/parking-lot-lld/components/entrance"
	"github/parking-lot-lld/components/exit"
	"github/parking-lot-lld/components/miscallaneous"
	parking_spot "github/parking-lot-lld/components/parking-spot"
	"github/parking-lot-lld/components/vehicles"
)

// Keep in mind this has to be a singleton class. So only one instance should be created at a time
type ParkingLot struct {
	id int
	name string
	address miscallaneous.Address
	parkingRate components.ParkingRate
	entrances map[int]entrance.Entrance
	exits map[int]exit.Exit
	parkingSpots map[string][]parking_spot.ParkingSpot
	tickets map[int]components.ParkingTicket
	ParkingLot *ParkingLot
}

// Implementing Singleton
func (pl ParkingLot) NewParkingLot (id int, name string, address miscallaneous.Address, parkingRate components.ParkingRate, entrances map[int]entrance.Entrance, exits map[int]exit.Exit, tickets map[int]components.ParkingTicket) *ParkingLot {
	if pl.ParkingLot != nil {
		return pl.ParkingLot
	} else {
		return &ParkingLot{
			id: id,
			name: name,
			address: address,
			parkingRate: parkingRate,
			entrances: entrances,
			exits: exits,
			tickets: tickets,
		}
	}
}

func (pl ParkingLot) AddEntrance(entrance entrance.Entrance) {
	pl.entrances[entrance.Id()] = entrance
}

func (pl ParkingLot) AddExit(exit exit.Exit) {
	pl.exits[exit.Id()] = exit
}

// Allows parking ticket to be received at multiple entrances
func (pl ParkingLot) GetParkingTicket(vehicle vehicles.Vehicle) *components.ParkingTicket {
	return nil
}

func (pl ParkingLot) IsFull(parkingSpotType string) bool {
	return len(pl.parkingSpots[parkingSpotType]) == 0
}