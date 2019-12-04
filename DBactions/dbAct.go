package DBactions

import (
	"database/sql"
	"fmt"
	"github.com/hubaxis/jwt-auth/Bike"
	"github.com/labstack/gommon/log"
)

type DBactions interface {
	Add(p *Bike.BikeItem) error
	Get(manufacturer string, model string) ([]Bike.BikeItem, error)
}

var count int

type dbAct struct {
	db *sql.DB
}

func (d dbAct) Add(p *Bike.BikeItem) error {
	if count == 0 {
		_, err := d.db.Exec("insert into allBikes (manufacturer, model, sizing, price, quantity) values ($1, $2, $3, $4, $5);", p.Manufacturer, p.Model, p.Size, p.Price, p.Quantity)
		count += 1
		return err
	}
	row := d.db.QueryRow("select * from allBikes where manufacturer = $1 and model = $2 and sizing = $3 and price = $4;", p.Manufacturer, p.Model, p.Size, p.Price)
	if row != nil {
		t := Bike.BikeItem{}
		err := row.Scan(&t.Id, &t.Manufacturer, &t.Model, &t.Size, &t.Price, &t.Quantity)
		_, err = d.db.Exec("update allBikes set quantity = $1 where id = $2;", t.Quantity+p.Quantity, t.Id)
		return err
	}
	return nil
}

func (d dbAct) Get(manufacturer string, model string) ([]Bike.BikeItem, error) {
	rows, err := d.db.Query("select * from allBikes where manufacturer = $1 and model = $2;", manufacturer, model)
	if err != nil {
		panic(err)
	}
	p := []Bike.BikeItem{}
	for rows.Next() {
		t := Bike.BikeItem{}
		err := rows.Scan(&t.Id, &t.Manufacturer, &t.Model, &t.Size, &t.Price, &t.Quantity)
		if err != nil {
			fmt.Println(err)
			continue
		}
		p = append(p, t)
	}
	return p, nil
}

func NewBikesDB() (DBactions, error) {
	db, err := sql.Open("sqlite3", "bikes.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE if not exists allBikes (id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		manufacturer  TEXT NOT NULL,
		model TEXT NOT NULL,
		sizing INTEGER NOT NULL,
		price INTEGER NOT  NULL, 
		quantity INTEGER DEFAULT 1)`)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow("SELECT COUNT(*) FROM allBikes").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return &dbAct{db: db}, nil
}
