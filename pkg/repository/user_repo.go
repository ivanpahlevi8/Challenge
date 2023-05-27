package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/authentication"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/configs"
	"github.com/ivanpahlevi8/synapsis_challange/pkg/model"

	_ "github.com/lib/pq"
)

// create variable for
var MyUser *UserRepo

// create object struct of item repo
type UserRepo struct {
	Config *configs.Config
}

// init repo
func InitUserRepo() *UserRepo {
	// create db connection based on credential of database
	MyUser = &UserRepo{}

	return MyUser
}

// create method to add data to database user item
func (item_repo *UserRepo) AddDataModel(modelItem model.UserAccount) (model.UserAccount, error) {
	// get variable from model
	getDataId := ""
	getDataUsername := modelItem.GetUsername()
	getDataPassword := modelItem.GetPassword()
	getDataFirstName := modelItem.GetFirstName()
	getDataLastName := modelItem.GetLastName()
	getDataAge := modelItem.GetAge()
	getDataListId := modelItem.GetListId()

	// casting password
	newPass, err := authentication.GenerateJWTToken(getDataPassword, getDataUsername)

	// check err
	if err != nil {
		fmt.Println("err : ", err.Error())
		// error happend
		errs := errors.New("error when casting password using jwt token : ")
		// error happen
		return model.UserAccount{}, errs
	}

	// set new pass
	getDataPassword = string(newPass)

	// casting id
	newId := uuid.New()

	// set id
	getDataId = newId.String()

	// create query for inserting data to database
	query := `INSERT INTO "usertable"("id", "username", "password", "first_name", "last_name", "age", "list_id") values($1, $2, $3, $4, $5, $6, $7)`

	// execute query with it's values
	_, err = item_repo.Config.DB.Exec(query, getDataId, getDataUsername, getDataPassword, getDataFirstName, getDataLastName, getDataAge, getDataListId)

	// check error
	if err != nil {
		// create error
		errs := errors.New("error when executing add query in user repository. check line 61 in user repository")
		// error happen
		return model.UserAccount{}, errs
	}

	return modelItem, nil
}

// create method to get data to user item database
func (item_repo *UserRepo) GetData(id string) (model.UserAccount, error) {
	// get data based on id

	// create variable of data
	var getData model.UserAccount

	// create query
	query := `SELECT * FROM "usertable" WHERE "id"=$1`

	// query from db
	dataCollect, err := item_repo.Config.DB.Query(query, id)

	// check for error
	if err != nil {
		// print error
		fmt.Println("error in db : ", err.Error())
		// error happen
		errs := errors.New("error happen when getting data using id")
		// return
		return model.UserAccount{}, errs
	}

	// iterate data query to find certain data with certain id
	for dataCollect.Next() {
		// create variable to a value from data query
		var getId string
		var getUsername string
		var getPassword string
		var getFirstName string
		var getLastName string
		var getAge int
		var getListId string

		// assign value to variable
		err := dataCollect.Scan(&getId, &getUsername, &getPassword, &getFirstName, &getLastName, &getAge, &getListId)

		// check error
		if err != nil {
			// create error message
			errs := errors.New("erorr happen when assign value from data query")
			// reutrn err
			return model.UserAccount{}, errs
		}

		// check if id correct
		if getId == id {
			// assign value to getData
			getData.SetId(getId)
			getData.SetUsername(getUsername)
			getData.SetPassword(getPassword)
			getData.SetFirstName(getFirstName)
			getData.SetLastName(getLastName)
			getData.SetAge(int32(getAge))
			getData.SetListId(getListId)

			// exit iterate
			break
		}
	}

	// if success
	return getData, nil
}

// create method to get all data
func (item_repo *UserRepo) GetAllData() ([]model.UserAccount, error) {
	// create slice of data
	var allData []model.UserAccount

	// create data object to hold object
	var dataObject model.UserAccount

	// create query to get all data
	query := `SELECT * FROM "usertable"`

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
		// create variable to hold value
		var getIdCheck string
		var getUsernameCheck string
		var getPasswordCheck string
		var getFirstNameCheck string
		var getLastNameCheck string
		var getAgeCheck int
		var getListId string

		// assign value to variable
		getAllData.Scan(&getIdCheck, &getUsernameCheck, &getPasswordCheck, getFirstNameCheck, &getLastNameCheck, &getAgeCheck, &getListId)

		// assign each value to object
		dataObject.SetId(getIdCheck)
		dataObject.SetUsername(getUsernameCheck)
		dataObject.SetPassword(getPasswordCheck)
		dataObject.SetFirstName(getFirstNameCheck)
		dataObject.SetLastName(getLastNameCheck)
		dataObject.SetAge(int32(getAgeCheck))
		dataObject.SetListId(getListId)

		// add data object to slice
		allData = append(allData, dataObject)
	}

	// return all data
	return allData, nil
}

// create method to update data
func (item_repo *UserRepo) UpdateData(newData model.UserAccount, id string) (model.UserAccount, error) {
	// update data based on id

	// get each column data from model
	getId := newData.GetId()
	getUsername := newData.GetUsername()
	getPassword := newData.GetPassword()
	getFirstName := newData.GetFirstName()
	getLastName := newData.GetLastName()
	getAge := newData.GetAge()
	getListId := newData.GetListId()

	// create query to get data from database with certain id
	// all data is used to check if data available or not in database
	query := `SELECT * FROM "usertable" where "id"=$1`
	allData, err := item_repo.Config.DB.Query(query, getId)

	// check error
	if err != nil {
		// error happen
		errs := errors.New("error happen when get all data using certain id in update data method")
		return model.UserAccount{}, errs
	}

	// create varible to check data
	var isDataAvail bool

	// iterate through data collect
	for allData.Next() {
		// create variable to hold data
		var getIdCheck string
		var getUsernameCheck string
		var getPasswordCheck string
		var getFirstNameCheck string
		var getLastNameCheck string
		var getAgeCheck int
		var getListIdCheck string

		// assign value to variable
		err := allData.Scan(&getIdCheck, &getUsernameCheck, &getPasswordCheck, &getFirstNameCheck, &getLastNameCheck, &getAgeCheck, &getListIdCheck)

		if err != nil {
			errs := errors.New("error when holding data in update method")
			return model.UserAccount{}, errs
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
		updateQuery := `UPDATE "usertable" SET "id" = $1, "username" = $2, "password" = $3, "first_name" = $4, "last_name" = $5, "age" = $6, "list_id" = $7 WHERE "id" = $8`

		// exec query
		_, err := item_repo.Config.DB.Query(updateQuery, getId, getUsername, getPassword, getFirstName, getLastName, getAge, getListId, getId)

		if err != nil {
			// error happen
			errs := errors.New("error when updating data from database")
			return model.UserAccount{}, errs
		}
	} else {
		// error happen
		errs := errors.New("data not valid to be updated in database")
		return model.UserAccount{}, errs
	}

	// return data if success
	return newData, nil
}

// create method to delete data in database
func (item_repo *UserRepo) DeleteData(id string) (model.UserAccount, error) {
	// create method to delete data from database

	// create variable to hold data as feedback
	var getData model.UserAccount

	// create query to get certain data with id as feedback
	query := `SELECT * FROM "usertable" WHERE "id"=$1`

	// start query prigress
	allData, err := item_repo.Config.DB.Query(query, id)

	// check error
	if err != nil {
		// error happend
		errs := errors.New("error happen when getting data with certain id in delete data")
		return model.UserAccount{}, errs
	}

	for allData.Next() {
		// create variable for holding data
		var getIdCheck string
		var getUsernameCheck string
		var getPasswordCheck string
		var getFirstNameCheck string
		var getLastNameCheck string
		var getAgeCheck int
		var getListIdCheck string

		// assign value to variable
		err := allData.Scan(&getIdCheck, &getUsernameCheck, &getPasswordCheck, &getFirstNameCheck, &getLastNameCheck, &getAgeCheck, &getListIdCheck)

		if err != nil {
			errs := errors.New("error when holding data in delete method")
			return model.UserAccount{}, errs
		}

		if getIdCheck == id {
			// assign value to data object
			getData.SetId(getIdCheck)
			getData.SetUsername(getUsernameCheck)
			getData.SetPassword(getPasswordCheck)
			getData.SetFirstName(getFirstNameCheck)
			getData.SetLastName(getLastNameCheck)
			getData.SetAge(int32(getAgeCheck))
			getData.SetListId(getListIdCheck)
			break
		}
	}

	// create dlete query
	deleteQuery := `DELETE FROM "usertable" WHERE "id"=$1`

	// start delete query
	_, err = item_repo.Config.DB.Query(deleteQuery, id)

	// check err
	if err != nil {
		// error happen
		errs := errors.New("error happen when running delete query in delete method")
		return model.UserAccount{}, errs
	}

	// return this if success
	return getData, nil
}
