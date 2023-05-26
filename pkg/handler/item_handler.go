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
var MyItemHandler *ItemHandler

// create  handler object
type ItemHandler struct {
	ItemService *service.ItemService
	Config      *configs.Config
}

// init function to create user handler
func InitItemHandler(item_service *service.ItemService) *ItemHandler {
	// create handler
	MyItemHandler = &ItemHandler{}

	MyItemHandler.ItemService = item_service

	return MyItemHandler
}

// create function to add data to database
func (item_handler *ItemHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	// set header for jason data
	w.Header().Set("Content-Type", "application/json")

	// create user data to hold value from body
	var newData model.ItemModel

	// encode from json body
	_ = json.NewDecoder(r.Body).Decode(&newData)

	// add new data to database
	getData, err := item_handler.ItemService.AddData(newData)

	// check error
	if err != nil {
		// if error happend
		fmt.Println("error happend : ", err.Error())
	}

	// send feedback as new data
	json.NewEncoder(w).Encode(getData)
}

// create function to get data based on id
func (item_handler *ItemHandler) GetDataById(w http.ResponseWriter, r *http.Request) {
	// create header for json data
	w.Header().Set("content-type", "application/json")

	// get id from url
	params := r.URL.Query()
	getId := params.Get("id")

	// get data from db
	getData, err := item_handler.ItemService.GetDataById(getId)

	// check err
	if err != nil {
		fmt.Println("error when query : ", err.Error())
		return
	}

	// encode data to json format
	json.NewEncoder(w).Encode(getData)
}

// create function to update data based on id
func (item_handler *ItemHandler) UpdateDataById(w http.ResponseWriter, r *http.Request) {
	// set header as json response
	w.Header().Set("Content-Type", "application/json")

	// get params input from uri
	params := r.URL.Query()

	// get id from params
	getId := params.Get("id")

	// create new data input to hold input body
	var newData model.ItemModel

	// assign data to input body
	_ = json.NewDecoder(r.Body).Decode(&newData)

	// update data
	getData, err := item_handler.ItemService.UpdateDataById(newData, getId)

	// check error
	if err != nil {
		// if error happend
		fmt.Println("error happend when updating data : ", err.Error())
	}

	// send feedback as json
	json.NewEncoder(w).Encode(getData)
}
