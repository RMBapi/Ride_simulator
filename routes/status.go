package routes

import (
	"net/http"
	"strconv"

	"example.com/Ride_simulator/models"
	"github.com/gin-gonic/gin"
)

func DriverStatusUpdate(context *gin.Context) {
	driverID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fatch driver id"})
		return
	}
	driverInfo, err := models.GetAllDriverStatusByID(driverID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch drivers status info"})
		return
	}

	var UpdatedStatus models.DriverStatus
	err = context.ShouldBindJSON(&UpdatedStatus)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse requested ststus"})
		return
	}

	UpdatedStatus.Id = driverID
	UpdatedStatus.PhoneNumber = driverInfo.PhoneNumber
	UpdatedStatus.Type = "driver"
	err = UpdatedStatus.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update driver status"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Driver status updated", "event": UpdatedStatus})

}
