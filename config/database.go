package config

import (
	"github.com/gowiki-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func DatabaseInit() {
	var err error

	var db *gorm.DB
	var userDb = goDotEnvVariable("USER_DB")
	var passwordDb = goDotEnvVariable("PASSWORD_DB")
	var portDb = goDotEnvVariable("PORT_DB")
	var nameDb = goDotEnvVariable("NAME_DB")

	db, err = gorm.Open(mysql.Open(userDb+":"+passwordDb+"@tcp(127.0.0.1:"+portDb+")/"+nameDb+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Comment{})
}
