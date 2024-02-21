package parking_spot

import "github/parking-lot-lld/components/vehicles"

type Large struct {
	id int
	isFree bool
	vehicle vehicles.Vehicle
}

func (m Large) NewLarge(id int, isFree bool) *Large {
	return &Large{
		id: id,
		isFree: isFree,
	}
}

func (m Large) IsFree() bool {
	return m.isFree
}

func (m Large) SetIsFree(isFree bool) {
	m.isFree = isFree
}

func (m Large) Id() int {
	return m.id
}

func (m Large) SetId(id int) {
	m.id = id
}

func (m Large) AssignVehicle(vehicle vehicles.Vehicle) bool {
	m.vehicle = vehicle
	return true
}

func (m Large) RemoveVehicle() bool {
	m.vehicle = nil
	return true
}