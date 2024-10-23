package utils

import "golang.org/x/crypto/bcrypt"


// Utility function for hashing password
func HashPassword(password string) (string, error){
	// Converts plain password to a hashed secure value
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	/*
	If the password is valid error = nil [true], else false
	*/
	return err == nil
}