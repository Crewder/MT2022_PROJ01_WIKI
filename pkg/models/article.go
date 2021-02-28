package models

import (
	"github.com/gowiki-api/pkg/tools"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

type Article struct {
	gorm.Model
	UserId  int
	User    User `gorm:"foreignKey:UserId"`
	Title   string
	Content string
	Slug    string
}

type Articles []Article

func init() {
	db.AutoMigrate(&Article{})
}

func GetAllArticles() []Article {
	var articles []Article
	db.Find(&articles)
	return articles
}

func GetArticleBySlug(slug string) *Article {
	var article Article
	db.Where("slug = ?", slug).Find(&article)
	return &article
}

func NewArticle(a *Article) {
	if a == nil {
		log.Fatal(a)
	}
	a.Slug = SlugUnique(strings.ToLower(tools.SanitizerSlug(a.Title)))
	db.Create(&a)
}

func UpdateArticle(a *Article) {
	a.Slug = SlugUnique(strings.ToLower(tools.SanitizerSlug(a.Title)))
	db.Save(&a)
}

func SlugUnique(title string) string {
	slugValid := false
	indexSlug := 1
	slug := strings.ToLower(tools.SanitizerSlug(title))

	for !slugValid {
		if GetArticleBySlug(slug).Slug != "" {
			slug = strings.ToLower(tools.SanitizerSlug(title)) + "-" + strconv.Itoa(indexSlug)
			indexSlug++
		} else {
			slugValid = true
		}
	}

	return slug
}