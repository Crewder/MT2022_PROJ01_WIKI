package models

import (
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	UserId    int     `json:"UserId" gorm:"not null"`
	User      User    `gorm:"foreignKey:UserId"`
	ArticleId int     `json:"ArticleId" gorm:"not null"`
	Article   Article `gorm:"foreignKey:ArticleId"`
	Comment   string  `json:"Comment" gorm:"not null"`
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
