package models

import (
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	UserId    int     `json:"UserId"`
	User      User    `gorm:"foreignKey:UserId"`
	ArticleId int     `json:"ArticleId"`
	Article   Article `gorm:"foreignKey:ArticleId"`
	Comment   string  `json:"Comment"`
}

type Comments []Comment

func init() {
	_ = db.AutoMigrate(&Comment{})
}

func NewComment(comment *Comment) {
	if comment == nil {
		log.Fatal(comment)
	}
	db.Create(&comment)
}

func GetAllCommentsByArticle(articleId string) []Comment {
	var comments []Comment
	db.Where("article_id = ?", articleId).Find(&comments)
	return comments
}
