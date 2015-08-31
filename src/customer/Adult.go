package customer
import (
	"room"
	"fee"
)

type Adult struct {

}

func (adult Adult) Charge(room room.RoomCharge) fee.Fee {
	return room.Charge()
}
