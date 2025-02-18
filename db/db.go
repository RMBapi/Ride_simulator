package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {

	DB, err = sql.Open("sqlite3", "Rides.db")

	if err != nil {
		panic("Couldn't connect database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {

	createRidersTable := `
	CREATE TABLE IF NOT EXISTS riders(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	phoneNumber TEXT NOT NULL UNIQUE
	)
	`
	_, err := DB.Exec(createRidersTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

	createDriversTable := `
	CREATE TABLE IF NOT EXISTS drivers(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	phoneNumber TEXT NOT NULL UNIQUE
	)
	`

	_, err = DB.Exec(createDriversTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

	createDriversStatusTable := `
	CREATE TABLE IF NOT EXISTS driverstatus(
	driver_id INTEGER,
	phoneNumber TEXT NOT NULL,
	status TEXT DEFAULT 'offline',
	FOREIGN KEY(driver_id) REFERENCES drivers(id)
	)
	`

	_, err = DB.Exec(createDriversStatusTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

	createTrigger := `
	CREATE TRIGGER IF NOT EXISTS after_driver_insert
	AFTER INSERT ON drivers
	BEGIN
		INSERT INTO driverstatus (driver_id, phoneNumber)
		VALUES (NEW.id, NEW.phoneNumber);
	END;
`

	_, err = DB.Exec(createTrigger)
	if err != nil {
		panic("Couldn't create trigger")
	}

	createRideRequestTable := `
	CREATE TABLE IF NOT EXISTS riderequest(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	driver_id INTEGER,
	rider_id INTEGER,
	status TEXT,
	FOREIGN KEY(driver_id) REFERENCES drivers(id),
	FOREIGN KEY(rider_id) REFERENCES riders(id)
	)
	`

	_, err = DB.Exec(createRideRequestTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

	completedRidesTable := `
	CREATE TABLE IF NOT EXISTS completedTrips(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	driver_id INTEGER,
	rider_id INTEGER,
	status TEXT,
	FOREIGN KEY(driver_id) REFERENCES drivers(id),
	FOREIGN KEY(rider_id) REFERENCES riders(id)
	)
	`

	_, err = DB.Exec(completedRidesTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

}
