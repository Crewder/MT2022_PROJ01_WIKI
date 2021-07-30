package storage

import (
	"github.com/gowiki-api/helpers"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

// GoDotEnvVariable
// Fetch variable from .env
func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func init() {
	var err error
	if db == nil {
		var userDb = GoDotEnvVariable("USER_DB")
		var passwordDb = GoDotEnvVariable("PASSWORD_DB")
		var portDb = GoDotEnvVariable("PORT_DB")
		var nameDb = GoDotEnvVariable("NAME_DB")
		db, err = gorm.Open(mysql.Open(userDb+":"+passwordDb+"@tcp(127.0.0.1:"+portDb+")/"+nameDb+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	}

	helpers.HandleError(500, err)
}

func GetDB() *gorm.DB {
	return db
}
