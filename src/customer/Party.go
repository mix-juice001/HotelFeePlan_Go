package customer

import (
	"fee"
	"room"
)

type Party struct {
}

var customers []Customer

func (party Party) Add(customer Customer) {
	customers = append(customers, customer)
}

func (party Party) Charge(room room.RoomCharge) fee.Fee {
	var result fee.Fee = fee.Fee{0}
	for _, customer := range customers {
        result = result.Add(customer.Charge(room))
	}
	return result
}
