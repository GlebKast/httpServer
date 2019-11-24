package main

import (
	_ "database/sql"
	"fmt"
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
	err = db.Add(tmp)
	if err != nil {
		panic(err)
	}

	tmp1, err := db.Get(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(tmp1.Model)

	//db, err := sql.Open("sqlite3", "bikes.db")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//_, err = db.Exec(`CREATE TABLE if not exists mtb (id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	//	manufacturer  TEXT NOT NULL,
	//	model TEXT NOT NULL,
	//	sizing INTEGER NOT NULL,
	//	price INTEGER NOT  NULL )`)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = db.Exec("insert into mtb (manufacturer, model, sizing, price) values ('Trek', 'Superfly', 19, 800)")
	//if err != nil{
	//	panic(err)
	//}
	//_, err = db.Exec("insert into mtb (manufacturer, model, sizing, price) values ('Trek', 'Slash', 21, 2000)")
	//if err != nil{
	//	panic(err)
	//}
	//result, err := db.Exec("insert into mtb (manufacturer, model, sizing, price) values ('Trek', 'Remedy', 20, 2200)")
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(result.RowsAffected())
	//
	//rows, err := db.Query("select * from mtb")
	//if err != nil {
	//	panic(err)
	//}
	//defer rows.Close()
	//bikeSl := []bikeItem{}
	//
	//for rows.Next(){
	//	p := bikeItem{}
	//	err := rows.Scan(&p.id, &p.manufacturer, &p.model, &p.size, &p.price)
	//	if err != nil{
	//		fmt.Println(err)
	//		continue
	//	}
	//	bikeSl = append(bikeSl, p)
	//}
	//for _, p := range bikeSl{
	//	fmt.Println(p.id, p.manufacturer, p.model, p.size, p.price)
	//}
}
