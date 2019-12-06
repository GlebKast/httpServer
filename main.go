package main

import (
	_ "database/sql"
	"github.com/hubaxis/jwt-auth/DBactions"
	"github.com/hubaxis/jwt-auth/handler"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//tests.StartDBTests()

	db, err := DBactions.NewBikesDB()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	bh := handler.New(db)
	e.GET("/Add", bh.AddBike)
	e.GET("/Get", bh.GetBike)
	e.GET("/Del", bh.DeleteBike)
	e.Logger.Fatal(e.Start(":8080"))
}
