package vehicles

import (
	"fmt"
	"github/parking-lot-lld/components"
)

type Car struct {
	licenseNo string
}

func (c *Car) NewCar(licenseNo string) *Car {
	return &Car{
		licenseNo: licenseNo,
	}
}

func (c *Car) AssignTicket(ticket components.ParkingTicket) {
	// Car's specific implementation of assignTicket
	fmt.Println("Car assigned ticket:", ticket)
}

func (c *Car) LicenseNo() string {
	return c.licenseNo
}

func (c *Car) SetLicenseNo(licenseNo string) {
	c.licenseNo = licenseNo
}