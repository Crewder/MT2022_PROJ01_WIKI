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

func init() {
	db.AutoMigrate(&Article{})
}

func GetAllArticles() []Article {
	var Articles []Article
	db.Find(&Articles)
	return Articles
}

func GetArticleById(Id int64) *Article {
	var article Article
	db.Where("ID = ?", Id).Find(&article)
	return &article
}
