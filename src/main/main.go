package main

import (
	"fmt"
	"customer"
	"room"
	"season"
)

func main() {
	var adult customer.Adult = customer.Adult{}
//	var twoRoom *room.TwoPeopleRoom = room.NewTwoPeopleRoom()
	twoRoom := room.NewTwoPeopleRoom(season.Busy{})
	twoRoom.Upgrade()

	fmt.Println(adult.Charge(twoRoom))


	var adult1 customer.Adult = customer.Adult{}
//	var adult2 *customer.Adult = new(customer.Adult)
	adult2 := new(customer.Adult)
	var child1 customer.Child = customer.Child{}
	var party customer.Party = customer.Party{}
	party.Add(adult1)
	party.Add(adult2)
	party.Add(child1)
	fmt.Println(party.Charge(twoRoom))
}
