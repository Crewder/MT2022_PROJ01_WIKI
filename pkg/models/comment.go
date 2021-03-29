package models

import (
	"gorm.io/gorm"
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

func NewComment(comment *Comment) bool {
	if comment == nil || comment.Comment == "" {
		return false
	}

	result := db.Create(&comment)

	if result.Error != nil {
		return false
	}

	return true
}

func GetAllCommentsByArticle(articleId string) []Comment {
	var comments []Comment
	db.Where("article_id = ?", articleId).Find(&comments)
	return comments
}

func GetComment(id string) *Comment {
	var comment Comment
	db.Where("id = ?", id).Find(&comment)
	return &comment
}

func DeleteComment(comment *Comment) {
	db.Delete(&comment)
}

func UpdateComment(comment *Comment) {
	db.Save(&comment)
}
