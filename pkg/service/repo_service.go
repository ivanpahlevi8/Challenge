package service

import (
	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/repository"
)

// create variable for user service
var MyShopService *ShopService

// crete struct object for service user
type ShopService struct {
	ShopRepo *repository.ItemRepo
	Config   *configs.Config
}

// init user service
func InitShopService(shop_repo *repository.ItemRepo) *ShopService {
	// create user service first
	MyShopService = &ShopService{}

	MyShopService.ShopRepo = shop_repo

	return MyShopService
}
