package models

import (
	"example.com/Ride_simulator/db"
)

type Driver struct {
	Id          int64
	PhoneNumber string `binding: "required`
	Type        string `binding: "required`
}

type DriverStatus struct {
	Status string
	Driver
}

func (u *Driver) Save() error {
	query := `INSERT INTO drivers(phoneNumber)
              VALUES(?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.PhoneNumber)
	if err != nil {
		return err
	}
	User_id, err := result.LastInsertId()
	u.Id = User_id
	return err

}

func GetAllDriverStatusByID(id int64) (*DriverStatus, error) {
	query := "SELECT * FROM driverstatus WHERE driver_id = ?"
	row := db.DB.QueryRow(query, id)
	var d DriverStatus
	err := row.Scan(&d.Id, &d.PhoneNumber, &d.Status)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (e *DriverStatus) Update() error {
	query := `
	UPDATE driverstatus
	SET status = ?
	WHERE driver_id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Status, e.Id)

	return err

}
