package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId    uint    `json:"UserId" gorm:"not null"`
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

func GetAllCommentsByArticle(articleId string) ([]Comment, bool) {
	var comments []Comment
	result := db.Where("article_id = ?", articleId).Find(&comments)
	if result.Error == nil {
		return comments, false
	}
	return comments, true
}

func GetComment(id string) (*Comment, bool) {
	var comment Comment
	result := db.Where("id = ?", id).Find(&comment)
	if result.Error == nil {
		return &comment, false
	}
	return &comment, true
}

func DeleteComment(comment *Comment) bool {
	result := db.Delete(&comment)
	if result.Error != nil {
		return false
	}
	return true
}

func UpdateComment(comment *Comment) bool {
	result := db.Save(&comment)
	if result.Error != nil {
		return false
	}
	return true
}
