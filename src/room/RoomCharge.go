package room
import "fee"

type RoomCharge interface {
	Charge() fee.Fee
	basicCharge() fee.Fee
}

