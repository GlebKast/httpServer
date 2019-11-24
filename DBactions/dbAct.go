package DBactions

import (
	"database/sql"
	"github.com/hubaxis/jwt-auth/Bike"
)

type DBactions interface {
	Add(bikeInfo Bike.BikeItem) error
	Get(myId int) (*Bike.BikeItem, error)
}

type dbAct struct {
	db *sql.DB
}

func (d dbAct) Add(bikeInfo Bike.BikeItem) error {
	_, err := d.db.Exec("insert into mtb (manufacturer, model, sizing, price) values ($1, $2, $3, $4);", bikeInfo.Manufacturer, bikeInfo.Model, bikeInfo.Size, bikeInfo.Price)
	return err
}

func (d dbAct) Get(myId int) (*Bike.BikeItem, error) {
	row := d.db.QueryRow("select * from mtb where id = $1", myId)
	p := Bike.BikeItem{}
	err := row.Scan(&p.Id, &p.Manufacturer, &p.Model, &p.Size, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func NewBikesDB() (DBactions, error) {
	db, err := sql.Open("sqlite3", "bikes.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE if not exists mtb (id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		manufacturer  TEXT NOT NULL,
		model TEXT NOT NULL,
		sizing INTEGER NOT NULL,
		price INTEGER NOT  NULL )`)
	if err != nil {
		return nil, err
	}
	return &dbAct{db: db}, nil
}
