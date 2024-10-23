package models

import (
	"api1/db"
	"api1/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// This will enable hashed password to be stored in database
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId

	return err
}

func (u User) ValidateCredentials() error {
	/*
	- Get the password of the user from the database with this particular e-mail
	- Get the id, it will be used for user authentication
	*/
	query := "SELECT id, password FROM users WHERE email = ?"

	/* 
	- Atleast one row will be found because we are look for a specific email, and in our database
	configuration—In the Users table, we made sure that the email column must contain unique values
	[email TEXT NOT NULL UNIQUE], so a single email address can't exist more than once on the table.

	- So we would get back one row or no results at all, if the email was not found.
	*/

	row := db.DB.QueryRow(query, u.Email)
	
	/*
	What row.Scan() does: It reads the values from a query result row and copies them into the provided variables,
	allowing you to work with the data directly.

	- id is read from the database and stored into the struct u.ID variable
	- password is read from the database and stored in the retrieved password variable

	This password[gotten from the database] will be compared with the one on the User struct[u.Password]
	*/
	var retrievedPassword string
	err := row.Scan(&u.ID ,&retrievedPassword)

	if err !=nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}