package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"starter-template/features/users"
	"starter-template/helper"
	"starter-template/helper/jwt"
)

type UserService struct {
	data users.UserDataInterface
	jwt  jwt.JWTInterface
}

func NewService(jwt jwt.JWTInterface, data users.UserDataInterface) *UserService {
	return &UserService{
		data: data,
		jwt:  jwt,
	}
}

func (u *UserService) Register(newData users.User) (*users.User, error) {
	_, err := u.data.GetByUsername(newData.Username)

	if err == nil {
		logrus.Error("Service : Username already registered")
		return nil, errors.New("username already registered by another user")
	}

	hashPassword, err := helper.HashPassword(newData.Password)
	if err != nil {
		logrus.Error("Service : Error Hash Password : ", err.Error())
		return nil, errors.New("hash password error")
	}

	newData.Password = hashPassword
	newData.IsAdmin = false
	newData.Status = true

	result, err := u.data.Register(newData)
	if err != nil {
		logrus.Error("Service : Error Register : ", err.Error())
		return nil, errors.New("error register")
	}

	return result, nil
}

func (u *UserService) Login(username string, password string) (*users.UserCredential, error) {
	dataUser, err := u.data.GetByUsername(username)

	if err != nil {
		logrus.Error("Service : Username Not Found")
		return nil, errors.New("username not found")
	}

	if err := helper.ComparePassword(dataUser.Password, password); err != nil {
		logrus.Error("Service : Password Incorrect")
		return nil, errors.New("password incorrect")
	}

	result, err := u.data.Login(username, password)

	if err != nil {
		logrus.Error("Service : Error Login : ", err.Error())
		return nil, errors.New("error process failed")
	}

	role := "user"

	if result.IsAdmin {
		role = "Admin"
	}

	tokenData := u.jwt.GenerateJWT(result.ID, result.Username, result.Email, result.PhoneNumber, role)

	if tokenData == nil {
		logrus.Error("Service : Generate Token Error : ", err.Error())
		return nil, errors.New("error token process failed")
	}

	response := new(users.UserCredential)
	response.Username = result.Username
	response.Access = tokenData

	return response, nil
}

func (u *UserService) Profile(id int) (*users.User, error) {
	res, err := u.data.GetByID(id)

	if err != nil {
		logrus.Error("Service : Get Profile Error : ", err.Error())
		return nil, errors.New("ERROR Get Profile")
	}

	return res, nil
}
