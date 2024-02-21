package parking_spot

import "github/parking-lot-lld/components/vehicles"

type ParkingSpot interface {
  IsFree() bool
  AssignVehicle(vehicle vehicles.Vehicle) bool
  RemoveVehicle() bool
}