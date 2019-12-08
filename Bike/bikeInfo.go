package Bike

import "fmt"

type BikeItem struct {
	Id           int
	Manufacturer string
	Model        string
	Size         int
	Price        int
	Quantity     int
}

type BikeDescription struct {
	Id           int
	Manufacturer string
	Model        string
	Description  string
}

func ShowBikeInfo(it BikeItem) {
	fmt.Print(it.Manufacturer + " " + it.Model + " ")
	fmt.Print(it.Price)
	fmt.Print(" ")
	fmt.Print(it.Size)
	fmt.Print(" ")
	fmt.Print(it.Quantity)
	fmt.Println()
}
