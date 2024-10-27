package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"starter-template/configs"
)

func InitDB(c *configs.ProgramConfig) *gorm.DB {
	// For Connection PostgreSQL
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.DBUser, c.DBPass, c.DBHost, c.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//	Un Comment For Connection MySQL
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Terjadi kesalahan pada database, error : ", err.Error())
		return nil
	}

	return db
}
