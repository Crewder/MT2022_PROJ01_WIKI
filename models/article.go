package models

import (
	"github.com/gowiki-api/tools"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Article struct {
	gorm.Model
	UserId  uint    `json:"UserId" gorm:"not null"`
	User    User    `gorm:"foreignKey:UserId"`
	Title   *string `json:"Title" gorm:"not null"`
	Content *string `json:"Content" gorm:"not null"`
	Slug    string  `json:"Slug" gorm:"not null"`
}

type Articles []Article

func init() {
	_ = db.AutoMigrate(&Article{})
}

func GetAllArticles() ([]Article, bool) {
	var articles []Article
	result := db.Find(&articles)
	if result.Error != nil {
		return articles, true
	}
	return articles, false
}

func GetArticleBySlug(slug string) (*Article, bool) {
	var article Article
	if len(slug) <= 0 {
		return &article, false
	}
	result := db.Where("slug = ?", slug).Find(&article)
	if article.Title == nil || article.Content == nil {
		return &article, true
	}
	if result.Error == nil {
		return &article, false
	}

	return &article, true
}

func GetArticleById(Id int64) (*Article, bool) {
	var article Article
	db.Where("ID = ?", Id).Find(&article)

	if article.Model.ID == 0 {
		return &article, false
	}
	return &article, true
}

func NewArticle(a *Article) bool {
	if a == nil || *a.Title == "" || *a.Content == "" {
		return false
	}
	if a.Title == nil {
		return false
	}
	a.Slug = SlugUnique(strings.ToLower(tools.SanitizerSlug(*a.Title)))
	result := db.Create(&a)
	if result.Error != nil {
		return false
	}
	return true
}

func UpdateArticle(a *Article) bool {
	if a == nil {
		return false
	}
	a.Slug = SlugUnique(strings.ToLower(tools.SanitizerSlug(*a.Title)))
	result := db.Save(&a)
	if result.Error != nil {
		return false
	}
	return true
}

func SlugUnique(title string) string {
	slugValid := false
	indexSlug := 1
	slug := strings.ToLower(tools.SanitizerSlug(title))

	for !slugValid {
		article, result := GetArticleBySlug(slug)
		if article.Slug != "" && result == true {
			slug = strings.ToLower(tools.SanitizerSlug(title)) + "-" + strconv.Itoa(indexSlug)
			indexSlug++
		} else if article.Slug == slug {
			slug = strings.ToLower(tools.SanitizerSlug(title)) + "-" + strconv.Itoa(indexSlug)
			indexSlug++
		} else {
			slugValid = true
		}
	}
	return slug
}

func DeleteArticle(article *Article) {
	db.Delete(&article)
}
