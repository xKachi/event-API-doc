package routes

import "github.com/gin-gonic/gin"


func RegisterRoutes(server *gin.Engine) {
	// handler for an incoming HTTP request - GET POST PUT PATCH DELETE
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // /events/1, /events/5
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	/// User Routes
	server.POST("/signup", signup)
	server.POST("/login", login)
} 