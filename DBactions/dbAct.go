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
	Del(manufacturer string, model string, size int, quantity int) error
	NewDescription(p *Bike.BikeDescription) bool
	GetDescription(manufacturer string, model string) (string, error)
	ShowAll() error
}

var count int

type DBAct struct {
	db *sql.DB
}

func (d DBAct) Add(p *Bike.BikeItem) error {
	row := d.db.QueryRow("select * from allBikes where manufacturer = $1 and model = $2 and sizing = $3 and price = $4;", p.Manufacturer, p.Model, p.Size, p.Price)
	t := Bike.BikeItem{}
	err := row.Scan(&t.Id, &t.Manufacturer, &t.Model, &t.Size, &t.Price, &t.Quantity)
	if count == 0 || err == sql.ErrNoRows {
		_, err := d.db.Exec("insert into allBikes (manufacturer, model, sizing, price, quantity) values ($1, $2, $3, $4, $5);", p.Manufacturer, p.Model, p.Size, p.Price, p.Quantity)
		count++
		return err
	}
	_, err = d.db.Exec("update allBikes set quantity = $1 where id = $2;", t.Quantity+p.Quantity, t.Id)
	count++
	return err
}

func (d DBAct) Get(manufacturer string, model string) ([]Bike.BikeItem, error) {
	rows, err := d.db.Query("select * from allBikes where manufacturer = $1 and model = $2;", manufacturer, model)
	defer rows.Close()
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

func (d DBAct) Del(manufacturer string, model string, size int, quantity int) error {
	if count != 0 {
		row := d.db.QueryRow("select * from allBikes where manufacturer = $1 and model = $2 and sizing = $3;", manufacturer, model, size)
		if row != nil {
			t := Bike.BikeItem{}
			err := row.Scan(&t.Id, &t.Manufacturer, &t.Model, &t.Size, &t.Price, &t.Quantity)
			if t.Quantity > quantity {
				_, err = d.db.Exec("update allBikes set quantity = $1 where id = $2;", t.Quantity-quantity, t.Id)
			} else {
				_, err = d.db.Exec("delete from allBikes where manufacturer = $1 and model = $2 and sizing = $3;", manufacturer, model, size)
			}
			return err
		}
		return nil
	}
	return nil
}

func (d DBAct) ShowAll() error {
	rows, err := d.db.Query("select * from allBikes;")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		t := Bike.BikeItem{}
		_ = rows.Scan(&t.Id, &t.Manufacturer, &t.Model, &t.Size, &t.Price, &t.Quantity)
		Bike.ShowBikeInfo(t)
	}
	return nil
}

func (d DBAct) NewDescription(p *Bike.BikeDescription) bool {
	res := d.db.QueryRow("select * from bikeDescriptions where manufacturer = $1 and model = $2;", p.Manufacturer, p.Model)
	err := res.Scan(&p.Id, &p.Manufacturer, &p.Model, &p.Description)
	if err != nil {
		_, err = d.db.Exec("insert into bikeDescriptions (manufacturer, model, description) values ($1, $2, $3);", p.Manufacturer, p.Model, p.Description)
		return true
	}
	return false
}

func (d DBAct) GetDescription(manufacturer string, model string) (string, error) {
	res := d.db.QueryRow("select * from bikeDescriptions where manufacturer = $1 and model = $2;", manufacturer, model)
	t := Bike.BikeDescription{}
	if res != nil {
		err := res.Scan(&t.Id, &t.Manufacturer, &t.Model, &t.Description)
		return t.Description, err
	}
	return "Sorry. No description available", nil
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
	_, err = db.Exec(`CREATE TABLE if not exists bikeDescriptions (id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		manufacturer  TEXT NOT NULL,
		model TEXT NOT NULL,
		description TEXT NOT NULL)`)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow("SELECT COUNT(*) FROM allBikes").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return &DBAct{db: db}, nil
}
