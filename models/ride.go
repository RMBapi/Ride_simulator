package models

import (
	"errors"

	"example.com/Ride_simulator/db"
)

type RideRequest struct {
	Id                int64
	RiderId           int64
	RiderPhoneNumber  string
	DriverId          int64
	DriverPhoneNumber string
	Status            string
}

func SendOnlineDriverlist() ([]int64, error) {
	query := `SELECT driver_id FROM driverstatus WHERE status = ?`
	rows, err := db.DB.Query(query, "online")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var driverIDs []int64

	for rows.Next() {
		var driverID int64
		err = rows.Scan(&driverID)
		if err != nil {
			return nil, err
		}
		driverIDs = append(driverIDs, driverID)
	}

	return driverIDs, nil

}

func InTripDiverList() ([]int64, error) {

	query := `SELECT driver_id FROM riderequest WHERE status = ?`
	rows, err := db.DB.Query(query, "start")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var driverIDs []int64
	for rows.Next() {
		var driverID int64
		err = rows.Scan(&driverID)
		if err != nil {
			return nil, err
		}
		driverIDs = append(driverIDs, driverID)
	}

	return driverIDs, nil

}

func EligableDriverID(list1, list2 []int64) int64 {

	var eligible int64 = 0

	for i := 0; i < len(list1); i++ {
		counter := 0
		for j := i; j < len(list2); j++ {
			if list1[i] == list2[j] {
				counter++
				break
			}
		}
		if counter == 0 {
			eligible = list1[i]
			break
		}

	}

	return eligible

}

func (e *RideRequest) Save() error {
	query := `
   INSERT INTO riderequest(driver_id,rider_id,status)
   VALUES(?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	e.Status = "start"
	result, err := stmt.Exec(e.DriverId, e.RiderId, e.Status)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id
	return err

}

func (e *RideRequest) IsDriverInTrip() (bool, int64, string, error) {
	query := `SELECT rider_id,status FROM riderequest WHERE driver_id = ?`
	row := db.DB.QueryRow(query, e.DriverId)
	var state string
	var rider int64
	err := row.Scan(&rider, &state)
	if err != nil {
		return false, 0, "", errors.New("couldn't read scan driver status")
	}

	return state == "start", rider, state, nil

}

func (e *RideRequest) IsUserInTrip() (bool, int64, string, error) {
	query := `SELECT driver_id,status FROM riderequest WHERE rider_id= ?`
	row := db.DB.QueryRow(query, e.RiderId)
	var state string
	var driver int64
	err := row.Scan(&driver, &state)
	if err != nil {
		return false, 0, "", errors.New("couldn't read scan driver status")
	}

	return state == "start", driver, state, nil

}

func (e *RideRequest) CompleteTrip() error {
	query := `
   INSERT INTO completedTrips(driver_id,rider_id,status)
   VALUES(?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	e.Status = "end"
	result, err := stmt.Exec(e.DriverId, e.RiderId, e.Status)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return errors.New("Couldn't get the last id")
	}

	e.Id = id
	err = updateTheRideRequestTable(e.DriverId, e.RiderId, "start")
	if err != nil {
		return err
	}
	return err
}

func updateTheRideRequestTable(DriverId int64, RiderId int64, status string) error {

	query := `DELETE FROM riderequest 
	          WHERE driver_id = ? AND rider_id = ? AND status = ?`

	_, err := db.DB.Exec(query, DriverId, RiderId, status)
	if err != nil {
		return errors.New("failed to execute delete query")
	}
	return nil

}
