package models

import (
	"gorm.io/gorm"
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
