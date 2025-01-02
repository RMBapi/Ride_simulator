package main

import (
	"example.com/Ride_simulator/db"
	"example.com/Ride_simulator/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
