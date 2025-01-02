package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/v1/riders", HandleRequest)
	server.GET("/api/v1/riders", HandleRequest)
	server.POST("/api/v1/drivers", registerDrivers)
	server.PUT("/drivers/:id/status", DriverStatusUpdate)
}
