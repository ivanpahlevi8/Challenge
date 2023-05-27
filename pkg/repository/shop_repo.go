package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/model"
	"github.com/lib/pq"
)

// create variale for item repo
var MyShop *ShopRepo

// create struct data object for item repo
type ShopRepo struct {
	Config *configs.Config
}

// create function to init Item Repo
func InitShopRepo() *ShopRepo {
	// create credensial of db table
	MyShop = &ShopRepo{}

	return MyShop
}

// create function to add shop data to database
func (shop_repo *ShopRepo) AddShopeItem(newShop model.ShopModel) (model.ShopModel, error) {
	// get variable from data
	getId := ""
	getAllItem := newShop.GetAllItems()

	// casting new id
	newId := uuid.New()
	getId = newId.String()

	// create query for insert data
	query := `INSERT INTO "shoptable"("id", "all_items") VALUES($1, $2)`

	// start query
	_, err := shop_repo.Config.DB.Exec(query, getId, pq.Array(getAllItem))

	// check error
	if err != nil {
		fmt.Println("error in shop repo : ", err.Error())
		// if error happend
		errors := errors.New("error when inserting item to database")
		return model.ShopModel{}, errors
	}

	// if query success
	return newShop, nil
}

// create function to get shop chart from database
func (shop_repo *ShopRepo) GetShop(id string) (model.ShopModel, error) {
	// get data based on id

	// create variable of data
	var getData model.ShopModel

	// create query
	query := `SELECT * FROM "shoptable" WHERE "id"=$1`

	// query from db
	dataCollect, err := shop_repo.Config.DB.Query(query, id)

	// check for error
	if err != nil {
		// print error
		fmt.Println("error in db : ", err.Error())
		// error happen
		errs := errors.New("error happen when getting data using id")
		// return
		return model.ShopModel{}, errs
	}

	// iterate data query to find certain data with certain id
	for dataCollect.Next() {
		// create variable to a value from data query
		var getId string
		var getAllItems []string

		// assign value to variable
		err := dataCollect.Scan(&getId, &getAllItems)

		// check error
		if err != nil {
			// create error message
			errs := errors.New("erorr happen when assign value from data query")
			// reutrn err
			return model.ShopModel{}, errs
		}

		// check if id correct
		if getId == id {
			// assign value to getData
			getData.SetId(getId)
			getData.SetAllItems(getAllItems)

			// exit iterate
			break
		}
	}

	// if success
	return getData, nil
}

// create function to update data in database
func (shop_repo *ShopRepo) UpdateShopData(newData model.ShopModel, id string) (model.ShopModel, error) {
	// update data based on id

	// get each column data from model
	getId := newData.GetId()
	getAllItems := newData.GetAllItems()

	// create query to get data from database with certain id
	// all data is used to check if data available or not in database
	query := `SELECT * FROM "shoptable" where "id"=$1`
	allData, err := shop_repo.Config.DB.Query(query, getId)

	// check error
	if err != nil {
		// error happen
		errs := errors.New("error happen when get all data using certain id in update data method")
		return model.ShopModel{}, errs
	}

	// create varible to check data
	var isDataAvail bool

	// iterate through data collect
	for allData.Next() {
		// create variable to hold data
		var getIdCheck string
		var getAllItemsCheck []string

		// assign value to variable
		err := allData.Scan(&getIdCheck, &getAllItemsCheck)

		if err != nil {
			errs := errors.New("error when holding data in update method")
			return model.ShopModel{}, errs
		}

		if getIdCheck == id {
			// if data exist
			isDataAvail = true
			break
		} else {
			isDataAvail = false
		}
	}

	// check if data is not null
	if isDataAvail {
		// if data available
		// create query for updating data
		updateQuery := `UPDATE "shoptable" SET "id" = $1, "all_items" = $2 WHERE "id" = $3`

		// exec query
		_, err := shop_repo.Config.DB.Query(updateQuery, getId, getAllItems, getId)

		if err != nil {
			// error happen
			errs := errors.New("error when updating data from database")
			return model.ShopModel{}, errs
		}
	} else {
		// error happen
		errs := errors.New("data not valid to be updated in database")
		return model.ShopModel{}, errs
	}

	// return data if success
	return newData, nil
}
