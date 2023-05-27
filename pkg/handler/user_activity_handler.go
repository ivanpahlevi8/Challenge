package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/service"
)

// create variable for user handler
var MyUserActivityHandler *UserActivityHandler

// create  handler object
type UserActivityHandler struct {
	UserService *service.UserService
	ItemService *service.ItemService
	ShopService *service.ShopService
	Config      *configs.Config
}

// init function to create user handler
func IntiUserActivityHandler(user_service *service.UserService) *UserActivityHandler {
	// create handler
	MyUserActivityHandler = &UserActivityHandler{}

	MyUserActivityHandler.UserService = user_service

	return MyUserActivityHandler
}

// create function for user to add item to chart
func (user_activity *UserActivityHandler) UserAddItem(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// get username from session in login activity
	getUserUsername := user_activity.Config.Session.GetString(r.Context(), "username")

	// create user obj
	user, err := user_activity.UserService.GetUserByUsername(getUserUsername)

	// check err
	if err != nil {
		fmt.Println("error in getting user based on username : ", err.Error())
	}

	// create params
	params := r.URL.Query()
	item_id := params.Get("item_id")
	item, err := user_activity.ItemService.GetDataById(item_id)

	// check err
	if err != nil {
		fmt.Println("error in getting item based on item id : ", err.Error())
	}

	// create slic
	var allItems []string

	// get slice from user
	shopModel, err := user_activity.ShopService.GetData(user.GetListId())

	// check err
	if err != nil {
		fmt.Println("error in user add item : ", err.Error())
	}

	// assighn slice
	allItems = shopModel.GetAllItems()

	// add current item
	allItems = append(allItems, item.GetItemName())

	// updating all item in shop model
	shopModel.SetAllItems(allItems)

	// update item quantity in store
	item.SetItemQuantity(item.GetItemQuantity() - 1)

	// updating item
	_, err = user_activity.ItemService.UpdateDataById(item, item.GetId())
	if err != nil {
		fmt.Println("err in update item : ", err)
		w.Write([]byte(fmt.Sprintf("err in update item : %s\n", err.Error())))
	}

	// updating shop model
	_, err = user_activity.ShopService.UpdateData(shopModel, shopModel.GetId())
	if err != nil {
		fmt.Println("err in update shop : ", err)
		w.Write([]byte(fmt.Sprintf("err in update shop : %s\n", err.Error())))
	}

	// give feedback to user
	w.Write([]byte("Success adding data in chart\n"))
}

// add function for user to remove item
func (user_activity *UserActivityHandler) UserRemoveItem(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// get username from session in login activity
	getUserUsername := user_activity.Config.Session.GetString(r.Context(), "username")

	// create user obj
	user, err := user_activity.UserService.GetUserByUsername(getUserUsername)

	// check err
	if err != nil {
		fmt.Println("error in getting user based on username : ", err.Error())
	}

	// create params and get item based on item id
	params := r.URL.Query()
	item_id := params.Get("item_id")
	item, err := user_activity.ItemService.GetDataById(item_id)

	// check err
	if err != nil {
		fmt.Println("error in getting item based on item id : ", err.Error())
	}

	// create slice
	var allItems []string

	// get slice from user
	shopModel, err := user_activity.ShopService.GetData(user.GetListId())

	// check err
	if err != nil {
		fmt.Println("error in user add item : ", err.Error())
	}

	// assighn slice
	allItems = shopModel.GetAllItems()

	/**
	DELETE ALGORITHM START HERE
	*/

	// create position variable for deleted item
	var position int

	// iterate through data
	for index, value := range allItems {
		if value == item.GetItemName() {
			position = index
			break
		}
	}

	// delete data
	allItems = append(allItems[:position], allItems[position+1:]...)

	// updating all item in shop model
	shopModel.SetAllItems(allItems)

	// update item quantity in store
	item.SetItemQuantity(item.GetItemQuantity() + 1)

	// updating item
	_, err = user_activity.ItemService.UpdateDataById(item, item.GetId())
	if err != nil {
		fmt.Println("err in update item : ", err)
		w.Write([]byte(fmt.Sprintf("err in update item : %s\n", err.Error())))
	}

	// updating shop model
	_, err = user_activity.ShopService.UpdateData(shopModel, shopModel.GetId())
	if err != nil {
		fmt.Println("err in update shop : ", err)
		w.Write([]byte(fmt.Sprintf("err in update shop : %s\n", err.Error())))
	}

	// give feedback to user
	w.Write([]byte("Success deleting data in chart \n"))
}

// create function for user to view all items in chart
func (user_activity *UserActivityHandler) UserViewAllItemInChart(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// get username from session in login activity
	getUserUsername := user_activity.Config.Session.GetString(r.Context(), "username")

	// create user obj
	user, err := user_activity.UserService.GetUserByUsername(getUserUsername)

	// check err
	if err != nil {
		fmt.Println("error in getting user based on username : ", err.Error())
	}

	// check err
	if err != nil {
		fmt.Println("error in getting item based on item id : ", err.Error())
	}

	// create slice
	var allItems []string

	// get slice from user
	shopModel, err := user_activity.ShopService.GetData(user.GetListId())

	// check err
	if err != nil {
		fmt.Println("error in user add item : ", err.Error())
	}

	// assighn slice
	allItems = shopModel.GetAllItems()

	if len(allItems) < 1 {
		w.Write([]byte("empty chart"))
	} else {
		json.NewEncoder(w).Encode(allItems)
	}

}
