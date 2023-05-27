package handler

import (
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

	fmt.Println("user list id : ", user.GetListId())

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
	}

	// updating shop model
	_, err = user_activity.ShopService.UpdateData(shopModel, shopModel.GetId())
	if err != nil {
		fmt.Println("err in update shop : ", err)
	}
}
