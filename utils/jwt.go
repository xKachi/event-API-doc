package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersceret"

func GenerateToken(email string, userId int64) (string, error) {

token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
	"email": email,
	"userId": userId,
	// Token valid 2hrs
	"exp": time.Now().Add(time.Hour * 2).Unix(),
})

	/*
		token return a complex value[*jwt.Token], so we need to get a single [string] that can be sent back to the 
		client, which will also be attached to future requests.

		- SignedString needs a key[secretKey] to be used for the signing process, and the key will later be used to verify
		incoming tokens
	*/

	return token.SignedString([]byte(secretKey))
}

// This function verifies the JWT token, so the user can access specific routes
func VerifyToken(token string) (int64, error) {
	/* The anonymous function automatically returns the secret key*/
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		/*Check if the token was verified with the signing method SHA256*/
		_, ok:= token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err !=nil {
		return 0, errors.New("Could not parse token.")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)


	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}