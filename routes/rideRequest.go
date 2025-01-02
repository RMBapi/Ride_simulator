package routes

import (
	"net/http"

	"example.com/Ride_simulator/models"
	"github.com/gin-gonic/gin"
)

func riderequest(context *gin.Context, r_id int64) {

	var ride models.RideRequest
	ride.RiderId = r_id
	riderInfo, err := models.GetUserByID(ride.RiderId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find rider info"})
		return
	}
	ride.RiderPhoneNumber = riderInfo.PhoneNumber

	onlineDriverList, err := models.SendOnlineDriverlist()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch online driver data"})
		return
	}

	intripDriverList, err := models.InTripDiverList()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch active ride driver data"})
		return
	}
	var eligableDriverID int64

	if len(intripDriverList) != 0 {
		eligableDriverID = models.EligableDriverID(onlineDriverList, intripDriverList)
	} else {
		eligableDriverID = onlineDriverList[0]
	}

	driverInfo, err := models.GetAllDriverStatusByID(eligableDriverID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch driver info"})
		return
	}

	ride.DriverId = eligableDriverID
	ride.DriverPhoneNumber = driverInfo.PhoneNumber

	err = ride.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could create ride"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "ride Created", "event": ride})

}

func endRide(context *gin.Context, D_id int64) {

	var ride models.RideRequest
	ride.DriverId = D_id

	driverInfo, err := models.GetAllDriverStatusByID(ride.DriverId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch driver info"})
		return
	}

	isOntrip, r_id, _, err := ride.IsDriverInTrip()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch driver status"})
		return
	}

	riderInfo, err := models.GetUserByID(r_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find rider info"})
		return
	}

	ride.RiderId = r_id
	ride.RiderPhoneNumber = riderInfo.PhoneNumber
	ride.DriverPhoneNumber = driverInfo.PhoneNumber

	if isOntrip == true {
		ride.CompleteTrip()
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Driver isn't in a trip"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "ride Ended", "event": ride})

}

func driverStatus(context *gin.Context, D_id int64) {

	var ride models.RideRequest

	ride.DriverId = D_id

	driverInfo, err := models.GetAllDriverStatusByID(ride.DriverId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch driver info"})
		return
	}

	if driverInfo.Status == "offline" {
		context.JSON(http.StatusOK, gin.H{"message": "Driver Status", "status": driverInfo.Status})
		return
	}

	isOntrip, r_id, state, _ := ride.IsDriverInTrip()

	if isOntrip {
		riderInfo, err := models.GetUserByID(r_id)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find rider info"})
			return
		}

		ride.RiderId = riderInfo.Id
		ride.RiderPhoneNumber = riderInfo.PhoneNumber
		ride.DriverPhoneNumber = driverInfo.PhoneNumber
		ride.Status = state

		context.JSON(http.StatusOK, gin.H{"message": "Driver Status", "event": ride})

	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Driver Status", "status": driverInfo.Status})
	}

}

func riderStatus(context *gin.Context, R_id int64) {

	var ride models.RideRequest

	ride.RiderId = R_id

	riderInfo, err := models.GetUserByID(ride.RiderId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find rider info"})
		return
	}

	isOntrip, d_id, state, _ := ride.IsUserInTrip()

	if isOntrip {
		ride.DriverId = d_id
		driverInfo, err := models.GetAllDriverStatusByID(ride.DriverId)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch driver info"})
			return
		}
		ride.RiderPhoneNumber = riderInfo.PhoneNumber
		ride.DriverPhoneNumber = driverInfo.PhoneNumber
		ride.Status = state

		context.JSON(http.StatusOK, gin.H{"message": "Rider Status", "event": ride})

	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Rider didn't in a trip"})
	}

}
