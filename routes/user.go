package routes

import (
	"net/http"

	"example.com/Ride_simulator/models"
	"github.com/gin-gonic/gin"
)

func registerRiders(context *gin.Context, p_n string) {
	var rider models.Rider

	b := models.IsValidPhoneNumber(p_n)

	if b != true {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalied Number"})
		return
	}
	rider.PhoneNumber = p_n

	err := rider.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create rider"})
		return
	}
	rider.Type = "rider"
	context.JSON(http.StatusCreated, gin.H{"message": "Rider Created", "event": rider})

}

func registerDrivers(context *gin.Context) {
	var driver models.Driver
	err := context.ShouldBindJSON(&driver)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse rider data"})
		return
	}

	err = driver.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create rider"})
		return
	}
	driver.Type = "driver"
	context.JSON(http.StatusCreated, gin.H{"message": "Driver Created", "event": driver})

}
