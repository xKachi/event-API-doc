package models

import (
	"api1/db"
	"time"
)

// Defines the shape of the event for the REST-API
type Event struct {
	ID int64 
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	// UserID links the event to the user that created the event
	UserID  int64
}

// Store slice of events
var events []Event = []Event{}

func (e Event) Save() error {

	// Adding the events to the database
	// The ? syntax is used here to prevent injection attacks
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	// get the last automatically generated ID(since it is auto-generated)
	id, err := result.LastInsertId()

	// Last generated id(int64 value), now used as the event id 
	e.ID = id

	return err

	// Managing events on a local slice
	//events = append(events, e)
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	// This method sends the query
	rows, err := db.DB.Query(query)
	
	if err !=nil {
		return nil, err
	}

	defer rows.Close()

	// loop through the rows to populate the event slice
	var events []Event

	for rows.Next() {
		// This struct is where the data being populated by the pointer will be stored for each row.
		var event Event

		// Pass a pointer to scan, so it is populated by the data from the row(one for every column).
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		// append the populated [event] to the events slice.
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		/*
		We have to return a pointer value for an event(*Event) this enable us return a nil value,
		reason being the null value for Event is Event{} not nil.
		*/
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
`

stmt, err := db.DB.Prepare(query)

if err != nil {
	return err
}

defer stmt.Close()

_ , err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	
	stmt, err := db.DB.Prepare(query)

	if err !=nil {
		return err
	}

	defer stmt.Close()

	_ , err = stmt.Exec(event.ID)

	return err

}