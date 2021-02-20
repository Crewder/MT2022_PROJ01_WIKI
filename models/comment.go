package models

import (
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	UserId    int
	User      User `gorm:"foreignKey:UserId"`
	ArticleId int
	Article   Article `gorm:"foreignKey:ArticleId"`
	Comment   string
}

type Comments []Comment

func init() {
	db.AutoMigrate(&Comment{})
}

func NewComment(comment *Comment) {
	if comment == nil {
		log.Fatal(comment)
	}
	db.Create(&comment)
}
