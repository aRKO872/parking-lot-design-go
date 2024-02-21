package vehicles

import (
	"fmt"
	"github/parking-lot-lld/components"
)

type MotorCycle struct {
	licenseNo string
}

func (c *MotorCycle) NewMotorcycle(licenseNo string) *MotorCycle {
	return &MotorCycle{
		licenseNo: licenseNo,
	}
}

func (c *MotorCycle) AssignTicket(ticket components.ParkingTicket) {
	// MotorCycle's specific implementation of assignTicket
	fmt.Println("MotorCycle assigned ticket:", ticket)
}

func (c *MotorCycle) LicenseNo() string {
	return c.licenseNo
}

func (c *MotorCycle) SetLicenseNo(licenseNo string) {
	c.licenseNo = licenseNo
}