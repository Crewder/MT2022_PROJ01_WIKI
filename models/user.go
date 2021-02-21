package models

import (
	"github.com/gowiki-api/config"
	"gorm.io/gorm"
	"log"
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

func NewUser(u *User) {
	db = config.GetDB()

	if u == nil {
		log.Fatal(u)
	}
	db.Create(&u)
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID = ?", Id).Find(&getUser)
	return &getUser, db
}

func GetUserByEmail(Email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("email = ?", Email).Find(&getUser)
	return &getUser, db
}
