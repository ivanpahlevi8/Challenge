package service

import (
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
