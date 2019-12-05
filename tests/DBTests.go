package tests

import (
	"fmt"
	"github.com/hubaxis/jwt-auth/Bike"
	"github.com/hubaxis/jwt-auth/DBactions"
)

func StartDBTests() {
	db, err := DBactions.NewBikesDB()
	if err != nil {
		panic(err)
	}
	tmp := Bike.BikeItem{}
	arr := []Bike.BikeItem{}
	fmt.Println("[Initial DB table]")
	err = db.ShowAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("[Adding trek superfly 5 19 2999 5]")
	tmp.Manufacturer = "trek"
	tmp.Model = "superfly 5"
	tmp.Size = 19
	tmp.Price = 2999
	tmp.Quantity = 5
	err = db.Add(&tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println("[Adding trek fuel ex 9.8 21 5000 1]")
	tmp.Manufacturer = "trek"
	tmp.Model = "fuel ex 9.8"
	tmp.Size = 21
	tmp.Price = 5000
	tmp.Quantity = 1
	err = db.Add(&tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println("[Adding trek fuel ex 9.8 18 4800 8]")
	tmp.Manufacturer = "trek"
	tmp.Model = "fuel ex 9.8"
	tmp.Size = 18
	tmp.Price = 4800
	tmp.Quantity = 8
	err = db.Add(&tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println("[Adding trek fuel ex 9.8 21 5000 3]")
	tmp.Manufacturer = "trek"
	tmp.Model = "fuel ex 9.8"
	tmp.Size = 21
	tmp.Price = 5000
	tmp.Quantity = 3
	err = db.Add(&tmp)
	if err != nil {
		panic(err)
	}
	fmt.Println("[DB table after adding items above]")
	err = db.ShowAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("[Deleting trek trek fuel ex 9.8 18 1]")
	err = db.Del("trek", "fuel ex 9.8", 18, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("[Getting trek trek fuel ex 9.8]")
	arr, err = db.Get("trek", "fuel ex 9.8")
	if err != nil {
		panic(err)
	}
	for _, tmp = range arr {
		Bike.ShowBikeInfo(tmp)
	}
	fmt.Println("[Deleting trek superfly 5 19 3]")
	err = db.Del("trek", "superfly 5", 19, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println("[Getting trek superfly 5]")
	arr, err = db.Get("trek", "superfly 5")
	if err != nil {
		panic(err)
	}
	for _, tmp = range arr {
		Bike.ShowBikeInfo(tmp)
	}
	fmt.Println("[DB table after all actions]")
	err = db.ShowAll()
	if err != nil {
		panic(err)
	}

}
