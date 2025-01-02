package models

import (
	"example.com/Ride_simulator/db"
)

type Rider struct {
	Id          int64
	PhoneNumber string `binding: "required`
	Type        string `binding: "required`
}

func (u *Rider) Save() error {
	query := `INSERT INTO riders(phoneNumber)
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

func GetUserByID(id int64) (*Rider, error) {
	query := "SELECT * FROM riders WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var d Rider
	err := row.Scan(&d.Id, &d.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
