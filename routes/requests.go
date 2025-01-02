package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleRequest(context *gin.Context) {
	if context.Request.Method == http.MethodPost {

		type requestBody struct {
			PhoneNumber string `json:"phoneNumber"`
			DriverID    *int64 `json:"driverId"`
			RiderID     *int64 `json:"riderId"`
		}

		var req requestBody
		err := context.ShouldBindJSON(&req)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse request"})
			return
		}
		if req.DriverID != nil {
			endRide(context, *req.DriverID)
			return
		} else if req.RiderID != nil {
			riderequest(context, *req.RiderID)
			return
		} else if req.PhoneNumber != "" {
			fmt.Print("hello2")
			registerRiders(context, req.PhoneNumber)
			return
		}
	} else if context.Request.Method == http.MethodGet {
		driverIDStr := context.Query("driver_id")
		riderIDStr := context.Query("rider_id")

		if driverIDStr != "" {
			driverID, err := strconv.ParseInt(driverIDStr, 10, 64)

			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid driver_id format"})
				return
			}
			driverStatus(context, driverID)
			return
		}

		if riderIDStr != "" {
			riderID, err := strconv.ParseInt(riderIDStr, 10, 64)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid rider_id format"})
				return
			}
			riderStatus(context, riderID)
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "Please provide either driver_id or rider_id"})
	}
}
