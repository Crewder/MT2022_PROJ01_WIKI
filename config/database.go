package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
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

	var userDb = goDotEnvVariable("USER_DB")
	var passwordDb = goDotEnvVariable("PASSWORD_DB")
	var portDb = goDotEnvVariable("PORT_DB")
	var nameDb = goDotEnvVariable("NAME_DB")

	db, err = gorm.Open(mysql.Open(userDb+":"+passwordDb+"@tcp(127.0.0.1:"+portDb+")/"+nameDb+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

}

func GetDB() *gorm.DB {
	return db
}
