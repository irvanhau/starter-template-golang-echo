package repository

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username    string `gorm:"column:username;type:varchar(255)"`
	Email       string `gorm:"column:email;type:varchar(255)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(255)"`
	Password    string `gorm:"column:password;type:varchar(255)"`
	IsAdmin     bool   `gorm:"column:is_admin;type:bool"`
	Status      bool   `gorm:"column:status;type:bool"`
}
