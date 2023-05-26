package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivanpahlevi8/synapsis_challange/pkg/authentication"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/model"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/service"
)

// create var for middleware
var MyMiddleware *MiddlewareObj

// create middleware object
type MiddlewareObj struct {
	MiddService *service.UserService
}

// init middleware
func InitMiddleware(service *service.UserService) {
	MyMiddleware = &MiddlewareObj{}

	MyMiddleware.MiddService = service
}

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func LoginMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// set header for jason data
			w.Header().Set("Content-Type", "application/json")

			// create model to hold data from body
			var loginModel model.LoginModel

			// parse dvalue from body
			json.NewDecoder(r.Body).Decode(&loginModel)

			// get username and password
			getPasswordInput := loginModel.GetPassword()
			getUsernameInput := loginModel.GetUsername()

			// get certain user by username
			getUser, err := MyMiddleware.MiddService.GetUserByUsername(getUsernameInput)

			// check error
			if err != nil {
				// error happen
				w.Write([]byte(err.Error()))
				return
			}

			// get password from user
			passwordToken := getUser.GetPassword()

			// check token and password input
			/**
			first input is passwod token that get from database
			second input is password from user input
			it return boolean, if true authetication success
			*/
			getUsername, err := authentication.ExtractClaims(passwordToken, getPasswordInput)

			// check error
			if err != nil {
				// error happen
				w.Write([]byte("Invalid user password"))
				return
			}

			fmt.Println("Status : ", getUsername)

			if getUsername {
				next.ServeHTTP(w, r)
			} else {
				w.Write([]byte("Wrong Password"))
			}
		})
}
