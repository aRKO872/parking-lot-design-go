package vehicles

import types "github/parking-lot-lld/components"

type Vehicle interface {
	LicenseNo() string
	SetLicenseNo(licenseNo string)
	AssignTicket(ticket types.ParkingTicket)
}