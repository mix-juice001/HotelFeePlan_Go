package room
import (
	"fee"
	"season"
)

type ThreePeopleRoom struct {
	*Room
}

const threeBasicCharge = 4000

func (self ThreePeopleRoom) Charge() fee.Fee {
	if self.upgraded {
		return self.BasicCharge()
	}
	return self.BasicCharge().Add(self.upgradeCharge())
}

func (self ThreePeopleRoom) BasicCharge() fee.Fee {
	return self.season.ChargeRoom4Three()
}

func NewThreePeopleRoom(season season.Season) *ThreePeopleRoom {
	return &ThreePeopleRoom{&Room{false, season}}
}
