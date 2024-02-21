package vehicles

import (
	"fmt"
	"github/parking-lot-lld/components"
)

type Truck struct {
	licenseNo string
}

func (c *Truck) NewTruck(licenseNo string) *Truck {
	return &Truck{
		licenseNo: licenseNo,
	}
}

func (c *Truck) AssignTicket(ticket components.ParkingTicket) {
	// Car's specific implementation of assignTicket
	fmt.Println("Truck assigned ticket:", ticket)
}

func (c *Truck) LicenseNo() string {
	return c.licenseNo
}

func (c *Truck) SetLicenseNo(licenseNo string) {
	c.licenseNo = licenseNo
}