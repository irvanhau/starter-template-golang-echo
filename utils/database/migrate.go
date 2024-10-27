package database

import (
	"gorm.io/gorm"
	repoUser "starter-template/features/users/repository"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(repoUser.User{})
}
