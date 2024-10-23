package routes

import (
	"net/http"
	"strconv"

	"api1/models"
	"api1/utils"

	"github.com/gin-gonic/gin"
)

// request handler for single events
func getEvent(context *gin.Context) {
	/*
	- get the id value encoded in the path
	- Here we are also converting the id gotten from the path to an integer from a string
	- We are also using http.StatusBadRequest response code, because if we fail to extract the id from the incoming 
	request path, it will seem as if an incorrect value hass been added to the path(since it can't be converted to an integer).
	*/

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// fetch event with id from the database
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)
}


// request handler function
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// Authorization token sent by the client
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Verify the token after extracting it from the request.
	/*
		This will ensure that only requests that have a valid token for the Authorization header
		will be able to create events
	*/
	userId ,err := utils.VerifyToken(token)

	if err !=nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// What is stored here adheres to the struct datatype
	var event models.Event

	/* This function is just like the Scan(), you pass a pointer
	of the variable that should be populated with data to it*/
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	} 
	
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	/*
	Get the event data from the database with this ID that we are updating
	*/
	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	var updatedEvent models.Event
	
	/* This function is just like the Scan(), you pass a pointer
	of the variable that should be populated with data to it.
	The event data being updated will also be populated here

	- Update the event data gotten from the database here
	*/
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	/*
	Since we are not creating a new event the ID should not be generated on the server, instead the
	existing event ID should be used  since it is only just being updated.
	*/
	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successful!"})
}


func deleteEvent(context *gin.Context) {
	// logic of the Id that should be deleted
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) 
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// logic for checking if the event with this Id already exists in the database
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}