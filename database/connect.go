package database

import (
	"../models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"os"
)

var DB *gorm.DB

func Connect() {

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +"@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = database
	
	database.AutoMigrate(&models.User{},&models.Role{} ,&models.Permission{}, &models.Product{} )
}
