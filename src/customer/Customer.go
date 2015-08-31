package customer

import (
	"fee"
	"room"
)

type Customer interface {
	Charge(room.RoomCharge) fee.Fee
}
