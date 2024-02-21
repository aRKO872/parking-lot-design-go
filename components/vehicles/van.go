package vehicles

import (
	"fmt"
	"github/parking-lot-lld/components"
)

type Van struct {
	licenseNo string
}

func (c *Van) NewVan(licenseNo string) *Van {
	return &Van{
		licenseNo: licenseNo,
	}
}

func (c *Van) AssignTicket(ticket components.ParkingTicket) {
	// Car's specific implementation of assignTicket
	fmt.Println("Van assigned ticket:", ticket)
}

func (c *Van) LicenseNo() string {
	return c.licenseNo
}

func (c *Van) SetLicenseNo(licenseNo string) {
	c.licenseNo = licenseNo
}