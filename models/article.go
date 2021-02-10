package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserId  int
	User    User `gorm:"foreignKey:UserId"`
	Title   string
	Content string
}

type Articles []Article
