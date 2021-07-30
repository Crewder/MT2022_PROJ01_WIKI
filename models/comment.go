package models

import (
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	UserId    uint    `json:"UserId" gorm:"not null"`
	User      User    `gorm:"foreignKey:UserId"`
	ArticleId *int    `json:"ArticleId" gorm:"not null"`
	Article   Article `gorm:"foreignKey:ArticleId"`
	Comment   *string `json:"Comment" gorm:"not null"`
}

type Comments []Comment

func init() {
	_ = db.AutoMigrate(&Comment{})
}

func NewComment(comment *Comment) (*Comment, error) {
	if comment == nil || *comment.Comment == "" {
		log.Fatal(400, "il faut saisir un commentaire ")
	}

	result := db.Create(&comment)
	return comment, result.Error
}

func GetAllCommentsByArticle(articleId string) ([]Comment, error) {
	var comments []Comment
	result := db.Where("article_id = ?", articleId).Find(&comments)

	return comments, result.Error
}

func GetComment(id string) (*Comment, error) {
	var comment Comment
	result := db.Where("id = ?", id).Find(&comment)

	return &comment, result.Error
}

func DeleteComment(comment *Comment) (*Comment, error) {
	result := db.Delete(&comment)

	return comment, result.Error
}
