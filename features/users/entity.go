package users

import "github.com/labstack/echo/v4"

type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"is_admin"`
	Status      bool   `json:"status"`
}

type UserCredential struct {
	Username string         `json:"username"`
	Access   map[string]any `json:"token"`
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(username string, password string) (*UserCredential, error)
	Profile(id int) (*User, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(username, password string) (*User, error)
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
}
