package seeds

import (
	"gorm.io/gorm"
	"starter-template/features/users"
	"starter-template/helper"
)

func CreateUser(db *gorm.DB, username, email, phone_number string) error {
	var countData int64
	db.Table("users").Where("username = ?", username).Count(&countData)

	if countData < 1 {
		hashPass, _ := helper.HashPassword("password")
		return db.Create(&users.User{
			Username:    username,
			Email:       email,
			PhoneNumber: phone_number,
			Password:    hashPass,
			IsAdmin:     true,
			Status:      true,
		}).Error
	}

	return nil
}
