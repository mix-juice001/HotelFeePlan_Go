package room
import
(
	"fee"
	"season"
)

type Room struct {
	upgraded bool
	season season.Season
}

const upgradeCharge = 1000

func (self Room) upgradeCharge() fee.Fee {
	return fee.Fee{upgradeCharge}
}

func (self Room) Upgrade() {
	self.upgraded = true
}
