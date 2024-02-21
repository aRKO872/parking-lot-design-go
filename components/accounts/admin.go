package accounts

import (
	display_board "github/parking-lot-lld/components/display-board"
	"github/parking-lot-lld/components/miscallaneous"
	parking_spot "github/parking-lot-lld/components/parking-spot"
	enum "github/parking-lot-lld/enums"
)

type Admin struct {
	id int
	username string
	status enum.AccountStatusEnum
	person miscallaneous.Person
}

func (a Admin) NewAdmin (id int, name string, status enum.AccountStatusEnum, person miscallaneous.Person) *Admin {
	return &Admin {
		id: id,
		username: name,
		status: status,
		person: person,
	}
}

func (a Admin) Login() string {
	return "token"
}

func (a Admin) Logout() bool {
	return true
}

func (a Admin) ResetPassword() string {
	return "token"
}

func (a Admin) AddEntry() bool {
	return true
}

func (a Admin) AddExit() bool {
	return true
}

func (a Admin) AddParkingSpot(floorName string, parkingSpot parking_spot.ParkingSpot) bool {
	return true
}

func (a Admin) AddDisplayBoard(floorName string, displayBoard display_board.DisplayBoard) bool {
	return true
}