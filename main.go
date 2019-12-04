package main

import (
	_ "database/sql"
	"github.com/hubaxis/jwt-auth/Bike"
	"github.com/hubaxis/jwt-auth/DBactions"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := DBactions.NewBikesDB()
	if err != nil {
		panic(err)
	}

	tmp := Bike.BikeItem{}
	tmp.Manufacturer = "trek"
	tmp.Model = "superfly 5"
	tmp.Size = 19
	tmp.Price = 2999
	tmp.Quantity = 1

	err = db.Add(&tmp)
	if err != nil {
		panic(err)
	}

	tmp.Manufacturer = "trek"
	tmp.Model = "superfly 5"
	tmp.Size = 19
	tmp.Price = 2999
	tmp.Quantity = 5

	err = db.Add(&tmp)
	if err != nil {
		panic(err)
	}

	arr := []Bike.BikeItem{}
	arr, err = db.Get("trek", "superfly 5")
	if err != nil {
		panic(err)
	}

	for _, tmp = range arr {
		Bike.ShowBikeInfo(tmp)
	}

}
