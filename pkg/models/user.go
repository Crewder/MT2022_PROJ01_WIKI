package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	ID       uint   `json:"ID"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Role     string `json:"Role"`
}

func init() {
	_ = db.AutoMigrate(&User{})
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

func NewUser(u *User) {
	if u == nil {
		log.Fatal(u)
	}
	db.Create(&u)
}

func GetUserByEmail(Email string) *User {
	var getUser User
	db := db.Where("email = ?", Email).Find(&getUser)
	if db.RowsAffected != 1 {
		log.Fatal(http.StatusBadRequest)
	}
	return &getUser
}

func PasswordIsValid(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
