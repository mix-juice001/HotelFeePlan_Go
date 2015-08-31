package season
import (
	"fee"
)

type Season interface {
	ChargeRoom4Two() fee.Fee
	ChargeRoom4Three() fee.Fee
}

type Low struct {

}

func (low Low) ChargeRoom4Two() fee.Fee {
	return fee.Fee{4000}
}

func (low Low) ChargeRoom4Three() fee.Fee {
	return fee.Fee{3500}
}

type Normal struct {

}

func (normal Normal) ChargeRoom4Two() fee.Fee {
	return fee.Fee{5000}
}

func (normal Normal) ChargeRoom4Three() fee.Fee {
	return fee.Fee{4500}
}

type Busy struct {

}

func (busy Busy) ChargeRoom4Two() fee.Fee {
	return fee.Fee{6000}
}

func (busy Busy) ChargeRoom4Three() fee.Fee {
	return fee.Fee{5500}
}
