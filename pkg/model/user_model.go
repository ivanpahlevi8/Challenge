package model

// create user model

type UserAccount struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int32  `json:"age"`
	ListId    string `json:"list_id"`
}

// create method for get and set id
func (account *UserAccount) GetId() string {
	return account.Id
}

func (account *UserAccount) SetId(setId string) {
	account.Id = setId
}

// create method for get and set username
func (account *UserAccount) GetUsername() string {
	return account.Username
}

func (account *UserAccount) SetUsername(username string) {
	account.Username = username
}

// create method for get and set password
func (account *UserAccount) GetPassword() string {
	return account.Password
}

func (account *UserAccount) SetPassword(password string) {
	account.Password = password
}

// create method for get and set First Name
func (account *UserAccount) GetFirstName() string {
	return account.FirstName
}

func (account *UserAccount) SetFirstName(firstName string) {
	account.FirstName = firstName
}

// create method for get and set Last Name
func (account *UserAccount) GetLastName() string {
	return account.LastName
}

func (account *UserAccount) SetLastName(lastName string) {
	account.LastName = lastName
}

// create method for get and set Age
func (account *UserAccount) GetAge() int32 {
	return account.Age
}

func (account *UserAccount) SetAge(age int32) {
	account.Age = age
}

// create method for get and set List Id
func (account *UserAccount) GetListId() string {
	return account.ListId
}

func (account *UserAccount) SetListId(listId string) {
	account.ListId = listId
}
