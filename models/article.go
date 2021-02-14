package models

import (
	"log"

	"github.com/gowiki-api/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Article struct {
	gorm.Model
	UserId  int
	User    User `gorm:"foreignKey:UserId"`
	Title   string
	Content string
}

type Articles []Article

func NewArticle(a *Article) {

	db = config.GetDB()

	if a == nil {
		log.Fatal(a)
	}

	db.Create(&a)

}
