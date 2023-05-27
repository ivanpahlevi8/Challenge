package service

import (
	"errors"

	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/model"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/repository"
)

// create variable for user service
var MyItemService *ItemService

// crete struct object for service user
type ItemService struct {
	ItemRepo *repository.ItemRepo
	Config   *configs.Config
}

// init user service
func InitItemService(user_repo *repository.ItemRepo) *ItemService {
	// create user service first
	MyItemService = &ItemService{}

	MyItemService.ItemRepo = user_repo

	return MyItemService
}

func (item_service *ItemService) AddData(newData model.ItemModel) (model.ItemModel, error) {
	// add data
	getData, err := item_service.ItemRepo.AddData(newData)

	// check error
	if err != nil {
		return model.ItemModel{}, err
	}

	// return data if success
	return getData, nil
}

func (item_service *ItemService) GetAllData() ([]model.ItemModel, error) {
	// create var
	var getAllItem []model.ItemModel
	var err error

	// get all data from repository
	getAllItem, err = item_service.ItemRepo.GetAllData()

	// check err
	if err != nil {
		// error happend
		errs := errors.New("error when getting all data from repo")
		return getAllItem, errs
	}

	// if success
	return getAllItem, nil
}

func (item_service *ItemService) GetDataById(id string) (model.ItemModel, error) {
	// get user
	getItem, err := item_service.ItemRepo.GetData(id)

	if err != nil {
		// if error happen
		return model.ItemModel{}, err
	}

	// return data if success
	return getItem, nil
}

func (item_service *ItemService) GetDataByItemName(itemName string) (model.ItemModel, error) {
	// get all user
	getAllItem, err := item_service.ItemRepo.GetAllData()

	// create variable to hold user id
	var itemId string

	// check error
	if err != nil {
		// error happen
		return model.ItemModel{}, err
	}

	// iterate all data in user
	for _, data := range getAllItem {
		// check data username
		if data.GetItemName() == itemName {
			// if data founded
			itemId = data.GetId()
			break
		} else {
			itemId = ""
		}
	}

	// check if data founded
	if itemId == "" {
		// if data not found
		errs := errors.New("no data with certain username")
		return model.ItemModel{}, errs
	}

	// get data based on user id
	getUserData, err := item_service.ItemRepo.GetData(itemId)

	// check err
	if err != nil {
		// error happen
		return model.ItemModel{}, err
	}

	// return data if success
	return getUserData, nil
}

func (item_service *ItemService) UpdateDataById(newData model.ItemModel, id string) (model.ItemModel, error) {
	// update user
	getItem, err := item_service.ItemRepo.UpdateData(newData, id)

	// check err
	if err != nil {
		return model.ItemModel{}, err
	}

	// if success
	return getItem, nil
}

func (item_service *ItemService) DeleteDataById(id string) (model.ItemModel, error) {
	// delete user
	getItem, err := item_service.ItemRepo.DeleteData(id)

	// check error
	if err != nil {
		// if error happen
		return model.ItemModel{}, err
	}

	// if success
	return getItem, nil
}
