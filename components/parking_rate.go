package components

type ParkingRate struct {
	hours float64
	ratePerHour float64
}

func (pr ParkingRate) NewParkingRate(hours float64, ratePerHour float64) *ParkingRate {
	return &ParkingRate{
		hours: hours,
		ratePerHour: ratePerHour,
	}
}

func (pr ParkingRate) Calculate() float64 {
	return pr.hours*pr.ratePerHour
}