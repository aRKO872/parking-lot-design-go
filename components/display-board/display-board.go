package display_board

import parking_spot "github/parking-lot-lld/components/parking-spot"

type DisplayBoard struct{
	id int
	spotMappings map[string][]parking_spot.ParkingSpot
}

func (db DisplayBoard) NewDisplayBoard(id int, spotMappings map[string][]parking_spot.ParkingSpot) *DisplayBoard {
	return &DisplayBoard{
		id: id,
		spotMappings: spotMappings,
	}
}

func (db DisplayBoard) AddParkingSpot(spotType string, spot parking_spot.ParkingSpot) {
	db.spotMappings[spotType] = append(db.spotMappings[spotType], spot)
}

func (db DisplayBoard) ShowFreeSlot() {
	// Show free Slot
}