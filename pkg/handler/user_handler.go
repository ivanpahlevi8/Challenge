package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/model"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/service"
)

// create variable for user handler
var MyUserHandler *UserHandler

// create  handler object
type UserHandler struct {
	UserService *service.UserService
	Config      *configs.Config
}

// init function to create user handler
func IntiUserHandler(user_service *service.UserService) *UserHandler {
	// create handler
	MyUserHandler = &UserHandler{}

	MyUserHandler.UserService = user_service

	return MyUserHandler
}

// create handler for login purposes
func (user_handler *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// set header for jason data
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Login success")

	w.Write([]byte("Login success"))
}

func (user_handler *UserHandler) AddDataDatabase(w http.ResponseWriter, r *http.Request) {
	// set header for jason data
	w.Header().Set("Content-Type", "application/json")

	// create user data to hold value from body
	var newData model.UserAccount

	// encode from json body
	_ = json.NewDecoder(r.Body).Decode(&newData)

	// add new data to database
	getData, err := user_handler.UserService.AddUser(newData)

	// check error
	if err != nil {
		// if error happend
		fmt.Println("error happend : ", err.Error())
	}

	// send feedback as new data
	json.NewEncoder(w).Encode(getData)
}

// create handler function to get data by id
func (user_handler *UserHandler) GetDataById(w http.ResponseWriter, r *http.Request) {
	// create header for json data
	w.Header().Set("content-type", "application/json")

	// get id from url
	params := r.URL.Query()
	getId := params.Get("id")

	// get data from db
	getData, err := user_handler.UserService.GetUserById(getId)

	// check err
	if err != nil {
		fmt.Println("error when query : ", err.Error())
		return
	}

	// encode data to json format
	json.NewEncoder(w).Encode(getData)
}

// create handler function to update data by id
func (user_handler *UserHandler) UpdateDataById(w http.ResponseWriter, r *http.Request) {
	// set header as json response
	w.Header().Set("Content-Type", "application/json")

	// get params input from uri
	params := r.URL.Query()

	// get id from params
	getId := params.Get("id")

	// create new data input to hold input body
	var newData model.UserAccount

	// assign data to input body
	_ = json.NewDecoder(r.Body).Decode(&newData)

	// update data
	getData, err := user_handler.UserService.UpdateUserById(newData, getId)

	// check error
	if err != nil {
		// if error happend
		fmt.Println("error happend when updating data : ", err.Error())
	}

	// send feedback as json
	json.NewEncoder(w).Encode(getData)
}

// create handler function to delete data by id
