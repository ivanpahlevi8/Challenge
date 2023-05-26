package model

// create object struct for login

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// create method to get and set username
func (login_model *LoginModel) GetUsername() string {
	return login_model.Username
}

func (login_model *LoginModel) SetUsername(username string) {
	login_model.Username = username
}

// create method for get and set password
func (login_model *LoginModel) GetPassword() string {
	return login_model.Password
}

func (login_model *LoginModel) SetPassword(password string) {
	login_model.Password = password
}
