package parking_spot

import "github/parking-lot-lld/components/vehicles"

type Compact struct {
	id int
	isFree bool
	vehicle vehicles.Vehicle
}

func (m Compact) NewCompact(id int, isFree bool) *Compact {
	return &Compact{
		id: id,
		isFree: isFree,
	}
}

func (m Compact) IsFree() bool {
	return m.isFree
}

func (m Compact) SetIsFree(isFree bool) {
	m.isFree = isFree
}

func (m Compact) Id() int {
	return m.id
}

func (m Compact) SetId(id int) {
	m.id = id
}

func (m Compact) AssignVehicle(vehicle vehicles.Vehicle) bool {
	m.vehicle = vehicle
	return true
}

func (m Compact) RemoveVehicle() bool {
	m.vehicle = nil
	return true
}