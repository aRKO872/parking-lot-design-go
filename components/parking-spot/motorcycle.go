package parking_spot

import "github/parking-lot-lld/components/vehicles"

type MotorCycle struct {
	id int
	isFree bool
	vehicle vehicles.Vehicle
}

func (m MotorCycle) NewMotorcycle(id int, isFree bool) *MotorCycle {
	return &MotorCycle{
		id: id,
		isFree: isFree,
	}
}

func (m MotorCycle) IsFree() bool {
	return m.isFree
}

func (m MotorCycle) SetIsFree(isFree bool) {
	m.isFree = isFree
}

func (m MotorCycle) Id() int {
	return m.id
}

func (m MotorCycle) SetId(id int) {
	m.id = id
}

func (m MotorCycle) AssignVehicle(vehicle vehicles.Vehicle) bool {
	m.vehicle = vehicle
	return true
}

func (m MotorCycle) RemoveVehicle() bool {
	m.vehicle = nil
	return true
}