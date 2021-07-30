package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	ID       uint   `json:"ID" gorm:"not null"`
	Name     string `json:"Name" gorm:"not null"`
	Email    string `json:"Email" gorm:"not null"`
	Password string `json:"Password" gorm:"not null"`
	Role     string `json:"Role"`
}

func init() {
	_ = db.AutoMigrate(&User{})
}

func GetAllUsers() ([]User, error) {
	var Users []User
	result := db.Find(&Users)

	return Users, result.Error
}

func GetUserById(Id int64) (*User, error) {
	var getUser User
	result := db.Where("ID = ?", Id).Find(&getUser)

	return &getUser, result.Error
}

func NewUser(u *User) (*User, error) {
	if u == nil {
		log.Fatal(u)
	}
	u.Role = "member"
	result := db.Create(&u)

	return u, result.Error
}

func GetUserByEmail(Email string) (*User, error) {
	var getUser User
	db := db.Where("email = ?", Email).Find(&getUser)
	if db.RowsAffected != 1 {
		log.Fatal(http.StatusBadRequest)
	}
	return &getUser, db.Error
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
