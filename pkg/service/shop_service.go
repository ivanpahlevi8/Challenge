package service

import (
	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/model"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/repository"
)

// create variable for user service
var MyShopService *ShopService

// crete struct object for service user
type ShopService struct {
	ShopRepo *repository.ShopRepo
	Config   *configs.Config
}

// init user service
func InitShopService(shop_repo *repository.ShopRepo) *ShopService {
	// create user service first
	MyShopService = &ShopService{}

	MyShopService.ShopRepo = shop_repo

	return MyShopService
}

// create function to add data to database
func (shop_service *ShopService) AddData(newData model.ShopModel) (model.ShopModel, error) {
	// add data
	getData, err := shop_service.ShopRepo.AddShopeItem(newData)

	// check error
	if err != nil {
		return model.ShopModel{}, err
	}

	// return data if success
	return getData, nil
}

// create function to get data from database
func (shop_service *ShopService) GetData(id string) (model.ShopModel, error) {
	// get user
	getItem, err := shop_service.ShopRepo.GetShop(id)

	if err != nil {
		// if error happen
		return model.ShopModel{}, err
	}

	// return data if success
	return getItem, nil
}

// create function to update data from database
func (shop_service *ShopService) UpdateData(newData model.ShopModel, id string) (model.ShopModel, error) {
	// update user
	getItem, err := shop_service.ShopRepo.UpdateShopData(newData, id)

	// check err
	if err != nil {
		return model.ShopModel{}, err
	}

	// if success
	return getItem, nil
}
