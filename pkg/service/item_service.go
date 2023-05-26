package service

import (
	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
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
