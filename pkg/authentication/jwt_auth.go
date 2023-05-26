package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// create function to generate jwt token
func GenerateJWTToken(message string, username string) (string, error) {
	// create byte variable
	secretMessage := []byte(message)

	// create token object
	token := jwt.New(jwt.SigningMethodHS256)

	// setupt token time
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(10 * time.Minute)

	claims["authorized"] = true

	claims["username"] = username

	// generate token
	generateToken, err := token.SignedString(secretMessage)

	// check error
	if err != nil {
		fmt.Println("error in jwt : ", err.Error())
		// create error
		errs := errors.New("error when generate jwt token based on message")
		return "", errs
	}

	// if success
	return generateToken, nil
}

func ExtractClaims(password string, realPassword string) (bool, error) {
	if password != "" {
		tokenString := password

		var keyfunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
			return []byte(realPassword), nil
		}

		token, err := jwt.Parse(tokenString, keyfunc)

		if err != nil {
			return false, err
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			return true, nil
		}
	}

	return false, nil
}
