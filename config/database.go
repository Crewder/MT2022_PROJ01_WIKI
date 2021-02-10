package config

import (
	"log"

	"github.com/gowiki-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseInit() {
	var err error

	var db *gorm.DB
	db, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:8889)/gowiki_api?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Comment{})
}
