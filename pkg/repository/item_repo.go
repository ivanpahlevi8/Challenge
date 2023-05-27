package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/model"
)

// create variale for item repo
var MyItem *ItemRepo

// create struct data object for item repo
type ItemRepo struct {
	Config *configs.Config
}

// create function to init Item Repo
func InitItemRepo() *ItemRepo {
	// create credensial of db table
	MyItem = &ItemRepo{}

	return MyItem
}

// create function add data to item repo
func (item_repo *ItemRepo) AddData(newData model.ItemModel) (model.ItemModel, error) {
	// get variable from data
	itemId := ""
	itemName := newData.GetItemName()
	itemCategory := newData.GetItemCategory()
	itemPrice := newData.GetItemPrice()
	itemQuantity := newData.GetItemQuantity()

	// casting new id
	newId := uuid.New()
	itemId = newId.String()

	// create query for insert data
	query := `INSERT INTO "itemtable"("id", "item_name", "item_category", "item_price", "item_quantity") VALUES($1, $2, $3, $4, $5)`

	// start query
	_, err := item_repo.Config.DB.Exec(query, itemId, itemName, itemCategory, itemPrice, itemQuantity)

	// check error
	if err != nil {
		// if error happend
		errors := errors.New("error when inserting item to database")
		return model.ItemModel{}, errors
	}

	// if query success
	return newData, nil
}

// create function go get data based on id
func (item_repo *ItemRepo) GetData(id string) (model.ItemModel, error) {
	// get data based on id

	// create variable of data
	var getData model.ItemModel

	// create query
	query := `SELECT * FROM "itemtable" WHERE "id"=$1`

	// query from db
	dataCollect, err := item_repo.Config.DB.Query(query, id)

	// check for error
	if err != nil {
		// print error
		fmt.Println("error in db : ", err.Error())
		// error happen
		errs := errors.New("error happen when getting data using id")
		// return
		return model.ItemModel{}, errs
	}

	// iterate data query to find certain data with certain id
	for dataCollect.Next() {
		// create variable to a value from data query
		var getId string
		var getItemName string
		var getItemCategory string
		var getItemPrice float64
		var getItemQuantity int32

		// assign value to variable
		err := dataCollect.Scan(&getId, &getItemName, &getItemCategory, &getItemPrice, &getItemQuantity)

		// check error
		if err != nil {
			// create error message
			errs := errors.New("erorr happen when assign value from data query")
			// reutrn err
			return model.ItemModel{}, errs
		}

		// check if id correct
		if getId == id {
			// assign value to getData
			getData.SetId(getId)
			getData.SetItemName(getItemName)
			getData.SetItemCategory(getItemCategory)
			getData.SetItemPrice(getItemPrice)
			getData.SetItemQuantity(getItemQuantity)

			// exit iterate
			break
		}
	}

	// if success
	return getData, nil
}

// get all data
func (item_repo *ItemRepo) GetAllData() ([]model.ItemModel, error) {
	// create slice of data
	var allData []model.ItemModel

	// create data object to hold object
	var dataObject model.ItemModel

	// create query to get all data
	query := `SELECT * FROM "itemtable"`

	// run query
	getAllData, err := item_repo.Config.DB.Query(query)

	// check error
	if err != nil {
		// error happend
		errs := errors.New("error when getting all data in get all data method")
		return allData, errs
	}

	// iterate all data
	for getAllData.Next() {
		var getIdCheck string
		var getItemNameCheck string
		var getItemCategoryCheck string
		var getItemPriceCheck float64
		var getItemQuantityCheck int32

		// assign value to variable
		err := getAllData.Scan(&getIdCheck, &getItemNameCheck, &getItemCategoryCheck, &getItemPriceCheck, &getItemQuantityCheck)

		if err != nil {
			errs := errors.New("error when holding data in update method")
			return allData, errs
		}

		// assign each value to object
		dataObject.SetId(getIdCheck)
		dataObject.SetItemName(getItemNameCheck)
		dataObject.SetItemCategory(getItemCategoryCheck)
		dataObject.SetItemPrice(getItemPriceCheck)
		dataObject.SetItemQuantity(getItemQuantityCheck)

		// add data object to slice
		allData = append(allData, dataObject)
	}

	// return all data
	return allData, nil
}

// create function to update from database
func (item_repo *ItemRepo) UpdateData(newData model.ItemModel, id string) (model.ItemModel, error) {
	// update data based on id

	// get each column data from model
	getId := newData.GetId()
	getItemName := newData.GetItemName()
	getItemCategory := newData.GetItemCategory()
	getItemPrice := newData.GetItemPrice()
	getItemQuantity := newData.GetItemQuantity()

	// create query to get data from database with certain id
	// all data is used to check if data available or not in database
	query := `SELECT * FROM "usertable" where "id"=$1`
	allData, err := item_repo.Config.DB.Query(query, getId)

	// check error
	if err != nil {
		// error happen
		errs := errors.New("error happen when get all data using certain id in update data method")
		return model.ItemModel{}, errs
	}

	// create varible to check data
	var isDataAvail bool

	// iterate through data collect
	for allData.Next() {
		// create variable to hold data
		var getIdCheck string
		var getItemNameCheck string
		var getItemCategoryCheck string
		var getItemPriceCheck float64
		var getItemQuantityCheck int32

		// assign value to variable
		err := allData.Scan(&getIdCheck, &getItemNameCheck, &getItemCategoryCheck, &getItemPriceCheck, &getItemQuantityCheck)

		if err != nil {
			errs := errors.New("error when holding data in update method")
			return model.ItemModel{}, errs
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
		updateQuery := `UPDATE "itemtable" SET "id" = $1, "item_name" = $2, "item_category" = $3, "item_price" = $4, "item_quantity" = $5 WHERE "id" = $6`

		// exec query
		_, err := item_repo.Config.DB.Query(updateQuery, getId, getItemName, getItemCategory, getItemPrice, getItemQuantity, getId)

		if err != nil {
			// error happen
			errs := errors.New("error when updating data from database")
			return model.ItemModel{}, errs
		}
	} else {
		// error happen
		errs := errors.New("data not valid to be updated in database")
		return model.ItemModel{}, errs
	}

	// return data if success
	return newData, nil
}

// create function to delete data in database
func (item_repo *ItemRepo) DeleteData(id string) (model.ItemModel, error) {
	// create method to delete data from database

	// create variable to hold data as feedback
	var getData model.ItemModel

	// create query to get certain data with id as feedback
	query := `SELECT * FROM "itemtable" WHERE "id"=$1`

	// start query prigress
	allData, err := item_repo.Config.DB.Query(query, id)

	// check error
	if err != nil {
		// error happend
		errs := errors.New("error happen when getting data with certain id in delete data")
		return model.ItemModel{}, errs
	}

	for allData.Next() {
		// create variable to hold data
		var getIdCheck string
		var getItemNameCheck string
		var getItemCategoryCheck string
		var getItemPriceCheck float64
		var getItemQuantityCheck int32

		// assign value to variable
		err := allData.Scan(&getIdCheck, &getItemNameCheck, &getItemCategoryCheck, &getItemPriceCheck, &getItemQuantityCheck)

		if err != nil {
			errs := errors.New("error when holding data in update method")
			return model.ItemModel{}, errs
		}

		// check if id correct
		if getIdCheck == id {
			// assign value to getData
			getData.SetId(getIdCheck)
			getData.SetItemName(getItemNameCheck)
			getData.SetItemCategory(getItemCategoryCheck)
			getData.SetItemPrice(getItemPriceCheck)
			getData.SetItemQuantity(getItemQuantityCheck)

			// exit iterate
			break
		}
	}

	// create dlete query
	deleteQuery := `DELETE FROM "itemtable" WHERE "id"=$1`

	// start delete query
	_, err = item_repo.Config.DB.Query(deleteQuery, id)

	// check err
	if err != nil {
		// error happen
		errs := errors.New("error happen when running delete query in delete method")
		return model.ItemModel{}, errs
	}

	// return this if success
	return getData, nil
}
