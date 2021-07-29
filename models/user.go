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

func Exists(id uint) bool {
	exist := false
	var getUser User
	anyUser := db.Where("ID = ?", id).Find(&getUser)

	if anyUser != nil {
		exist = true
	}
	return exist
}

func GetAllUsers() ([]User, bool) {
	var Users []User
	result := db.Find(&Users)

	if result.Error != nil {
		return Users, false
	}
	return Users, true
}

func GetUserById(Id int64) (*User, bool) {
	var getUser User
	result := db.Where("ID = ?", Id).Find(&getUser)
	if result.Error == nil {
		return &getUser, false
	}
	return &getUser, true
}

func NewUser(u *User) bool {
	if u == nil {
		log.Fatal(u)
	}
	u.Role = "member"
	result := db.Create(&u)
	if result.Error != nil {
		return false
	}
	return true
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
