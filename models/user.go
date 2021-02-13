package models

import (
	"github.com/gowiki-api/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Email    string
	Password string
}

func init() {
	config.DatabaseInit()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) *User {
	var getUser User
	db.Where("ID = ?", Id).Find(&getUser)
	return &getUser
}
