package handler

import (
	"github.com/hubaxis/jwt-auth/Bike"
	"github.com/hubaxis/jwt-auth/DBactions"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type BikeHandler struct {
	rep DBactions.DBactions
}

func New(rep DBactions.DBactions) *BikeHandler {
	return &BikeHandler{rep: rep}
}

func (u BikeHandler) AddBike(c echo.Context) error {
	bikeManufacturer := c.QueryParam("manufacturer")
	bikeModel := c.QueryParam("model")
	bikeSize := c.QueryParam("size")
	bikePrice := c.QueryParam("price")
	bikeQuantity := c.QueryParam("quantity")
	bS, err := strconv.Atoi(bikeSize)
	if err != nil {
		return err
	}
	bP, err := strconv.Atoi(bikePrice)
	if err != nil {
		return err
	}
	bQ, err := strconv.Atoi(bikeQuantity)
	if err != nil {
		return err
	}
	newItems := &Bike.BikeItem{Manufacturer: bikeManufacturer, Model: bikeModel, Size: bS, Price: bP, Quantity: bQ}
	err = u.rep.Add(newItems)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Added: "+newItems.Manufacturer+" "+newItems.Model+"; size: "+bikeSize+"; price: "+bikePrice+"; quantity: "+bikeQuantity+".")
}

func (u BikeHandler) GetBike(c echo.Context) error {
	bikeManufacturer := c.QueryParam("manufacturer")
	bikeModel := c.QueryParam("model")
	arr := []Bike.BikeItem{}
	arr, err := u.rep.Get(bikeManufacturer, bikeModel)
	if err != nil {
		return err
	}
	t := Bike.BikeItem{}
	var res string
	for _, t = range arr {
		res += t.Manufacturer + " " + t.Model + " " + strconv.Itoa(t.Size) + " " + strconv.Itoa(t.Price) + " " + strconv.Itoa(t.Quantity) + "\n"
	}
	return c.String(http.StatusOK, res)
}

func (u BikeHandler) DeleteBike(c echo.Context) error {
	bikeManufacturer := c.QueryParam("manufacturer")
	bikeModel := c.QueryParam("model")
	bikeSize := c.QueryParam("size")
	bikeQuantity := c.QueryParam("quantity")
	bS, err := strconv.Atoi(bikeSize)
	if err != nil {
		return err
	}
	bQ, err := strconv.Atoi(bikeQuantity)
	if err != nil {
		return err
	}
	err = u.rep.Del(bikeManufacturer, bikeModel, bS, bQ)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "Deleted: "+bikeManufacturer+" "+bikeModel+"; size: "+bikeSize+"; quantity: "+bikeQuantity+".")
}
