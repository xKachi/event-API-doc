package main

import (
	"api1/db"
	"api1/routes"

	"github.com/gin-gonic/gin"
)


func main () {
	// database instance
	db.InitDB()
	// server instance
	server := gin.Default()

	/*
	- register routes 
	- The server value passed to RegisterRoutes is a pointer because the Default() function from the
		gin package returns a pointer to the gin Engine [gin.Default → *gin.Engine]
	*/ 
	routes.RegisterRoutes(server)
	

	server.Run(":8080") // localhost:8080
}



