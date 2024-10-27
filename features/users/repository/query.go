package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"starter-template/features/users"
)

type UserData struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserData {
	return &UserData{
		db: db,
	}
}

func (ud *UserData) Register(newData users.User) (*users.User, error) {
	var dbData = new(User)
	dbData.Username = newData.Username
	dbData.Email = newData.Email
	dbData.PhoneNumber = newData.PhoneNumber
	dbData.Password = newData.Password
	dbData.IsAdmin = newData.IsAdmin
	dbData.Status = newData.Status

	if err := ud.db.Create(dbData).Error; err != nil {
		logrus.Error("REPOSITORY : Register Error : ", err.Error())
		return nil, err
	}

	return &newData, nil
}

func (ud *UserData) Login(username string, password string) (*users.User, error) {
	var dbData = new(User)
	dbData.Username = username
	dbData.Status = true

	var qry = ud.db.Where("username = ? AND status = ?", dbData.Username, dbData.Status).First(dbData)

	var dataCount int64
	qry.Count(&dataCount)

	if dataCount == 0 {
		logrus.Error("REPOSITORY : Login Error : User Not Found")
		return nil, errors.New("user not found")
	}

	if err := qry.Error; err != nil {
		logrus.Error("REPOSITORY : Login Error : ", err.Error())
		return nil, err
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Username = dbData.Username
	result.Email = dbData.Email
	result.PhoneNumber = dbData.PhoneNumber
	result.IsAdmin = dbData.IsAdmin
	result.Status = dbData.Status

	return result, nil
}

func (ud *UserData) GetByID(id int) (*users.User, error) {
	var dbData = new(User)
	dbData.ID = uint(id)
	dbData.Status = true
	
	var qry = ud.db.Where("id = ? AND status = ?", dbData.ID, dbData.Status).First(&dbData)

	var result = new(users.User)
	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By ID : ", err.Error())
		return nil, err
	}

	result.ID = dbData.ID
	result.Username = dbData.Username
	result.Email = dbData.Email
	result.PhoneNumber = dbData.PhoneNumber
	result.IsAdmin = dbData.IsAdmin
	result.Status = dbData.Status

	return result, nil
}

func (ud *UserData) GetByUsername(username string) (*users.User, error) {
	var dbData = new(User)
	dbData.Username = username
	dbData.Status = true

	var qry = ud.db.Where("username = ? AND status = ?", dbData.Username, dbData.Status).First(dbData)

	if err := qry.Error; err != nil {
		logrus.Error("DATA : Error Get By Username : ", err.Error())
		return nil, err
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Username = dbData.Username
	result.Email = dbData.Email
	result.PhoneNumber = dbData.PhoneNumber
	result.IsAdmin = dbData.IsAdmin
	result.Status = dbData.Status

	return result, nil
}
