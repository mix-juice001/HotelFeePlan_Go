package customer

import (
	"room"
	"fee"
)

type Child struct {

}

const chileRate = 0.5

func (child Child) Charge(room room.RoomCharge) fee.Fee {
	return room.Charge().Multiple(fee.Rate{chileRate})
}
