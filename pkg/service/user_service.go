package service

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/model"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/repository"
)

// create variable for user service
var MyUserService *UserService

// crete struct object for service user
type UserService struct {
	UserRepo  *repository.UserRepo
	ItemsRepo *repository.ShopRepo
	Config    *configs.Config
}

// init user service
func InitUserService(user_repo *repository.UserRepo) *UserService {
	// create user service first
	MyUserService = &UserService{}

	MyUserService.UserRepo = user_repo

	return MyUserService
}

/**
method for user service start here
*/

// create method to add user data
func (user_service *UserService) AddUser(userAdd model.UserAccount) (model.UserAccount, error) {
	// create shop data first
	var itemData model.ShopModel

	// set item data
	itemData.Id = uuid.New().String()

	// set item slice
	itemData.AllItems = []string{}

	// change user id based on item data
	userAdd.SetListId(itemData.GetId())

	// add data
	getData, err := user_service.UserRepo.AddDataModel(userAdd)

	// check error
	if err != nil {
		return model.UserAccount{}, err
	}

	// add shop
	_, err = user_service.ItemsRepo.AddShopeItem(itemData)

	if err != nil {
		log.Println(err.Error())
	}

	// return data if success
	return getData, nil
}

// create method to get user data based on id
func (user_service *UserService) GetUserById(id string) (model.UserAccount, error) {
	// get user
	getUser, err := user_service.UserRepo.GetData(id)

	if err != nil {
		// if error happen
		return getUser, err
	}

	// return data if success
	return getUser, nil
}

// create method to get user data based on username
func (user_service *UserService) GetUserByUsername(username string) (model.UserAccount, error) {
	// get all user
	getAllUser, err := user_service.UserRepo.GetAllData()

	// create variable to hold user id
	var userId string

	// check error
	if err != nil {
		// error happen
		return model.UserAccount{}, err
	}

	// iterate all data in user
	for _, data := range getAllUser {
		// check data username
		if data.GetUsername() == username {
			// if data founded
			userId = data.GetId()
			break
		} else {
			userId = ""
		}
	}

	// check if data founded
	if userId == "" {
		// if data not found
		errs := errors.New("no data with certain username")
		return model.UserAccount{}, errs
	}

	// get data based on user id
	getUserData, err := user_service.UserRepo.GetData(userId)

	// check err
	if err != nil {
		// error happen
		return model.UserAccount{}, err
	}

	// return data if success
	return getUserData, nil
}

// create method to get user data based on First Name
func (user_service *UserService) GetUserByFirstName(firstName string) (model.UserAccount, error) {
	// get all user
	getAllUser, err := user_service.UserRepo.GetAllData()

	// create variable to hold user id
	var userId string

	// check error
	if err != nil {
		// error happen
		return model.UserAccount{}, err
	}

	// iterate all data in user
	for _, data := range getAllUser {
		// check data username
		if data.GetFirstName() == firstName {
			// if data founded
			userId = data.GetId()
			break
		} else {
			userId = ""
		}
	}

	// check if data founded
	if userId == "" {
		// if data not found
		errs := errors.New("no data with certain first name")
		return model.UserAccount{}, errs
	}

	// get data based on user id
	getUserData, err := user_service.UserRepo.GetData(userId)

	// check err
	if err != nil {
		// error happen
		return model.UserAccount{}, err
	}

	// return data if success
	return getUserData, nil
}

// create method to update data in database
func (user_service *UserService) UpdateUserById(newUser model.UserAccount, id string) (model.UserAccount, error) {
	// update user
	getUser, err := user_service.UserRepo.UpdateData(newUser, id)

	// check err
	if err != nil {
		return model.UserAccount{}, err
	}

	// if success
	return getUser, nil
}

// create method to delete data in database
func (user_service *UserService) DeleteUserById(id string) (model.UserAccount, error) {
	// delete user
	getUser, err := user_service.UserRepo.DeleteData(id)

	// check error
	if err != nil {
		// if error happen
		return model.UserAccount{}, err
	}

	// if success
	return getUser, nil
}
