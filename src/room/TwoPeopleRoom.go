package room
import (
	"fee"
	"season"
)

type TwoPeopleRoom struct {
	*Room
}

const twoBasicCharge = 3000

func (self TwoPeopleRoom) Charge() fee.Fee {
	if self.upgraded {
		return self.basicCharge()
	}
	return self.basicCharge().Add(self.upgradeCharge())
}

func (self TwoPeopleRoom) basicCharge() fee.Fee {
	return self.season.ChargeRoom4Two()
}

func NewTwoPeopleRoom(season season.Season) *TwoPeopleRoom {
	return &TwoPeopleRoom{&Room{false, season}}
}
