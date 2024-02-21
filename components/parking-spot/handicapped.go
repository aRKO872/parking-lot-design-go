package parking_spot

import "github/parking-lot-lld/components/vehicles"

type Handicapped struct {
	id int
	isFree bool
	vehicle vehicles.Vehicle
}

func (m Handicapped) NewHandicapped(id int, isFree bool) *Handicapped {
	return &Handicapped{
		id: id,
		isFree: isFree,
	}
}

func (m Handicapped) IsFree() bool {
	return m.isFree
}

func (m Handicapped) SetIsFree(isFree bool) {
	m.isFree = isFree
}

func (m Handicapped) Id() int {
	return m.id
}

func (m Handicapped) SetId(id int) {
	m.id = id
}

func (m Handicapped) AssignVehicle(vehicle vehicles.Vehicle) bool {
	m.vehicle = vehicle
	return true
}

func (m Handicapped) RemoveVehicle() bool {
	m.vehicle = nil
	return true
}