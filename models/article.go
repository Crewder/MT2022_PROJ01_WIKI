package models

import (
	"github.com/gowiki-api/tools"
	"gorm.io/gorm"
	"log"
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

func GetAllArticles() ([]Article, error) {
	var articles []Article
	result := db.Find(&articles)
	return articles, result.Error
}

func GetArticleBySlug(slug string) (*Article, error) {
	var article Article
	if len(slug) <= 0 {
		log.Fatal(400, "il faut saisir le slug")
	}
	result := db.Where("slug = ?", slug).Find(&article)
	if article.Title == nil || article.Content == nil {
		return &article, result.Error
	}

	return &article, result.Error
}

func GetArticleById(Id int64) (*Article, error) {
	var article Article
	result := db.Where("ID = ?", Id).Find(&article)

	if article.Model.ID == 0 {
		return &article, result.Error
	}
	return &article, result.Error
}

func NewArticle(a *Article) (*Article, error) {
	if a == nil || *a.Title == "" || *a.Content == "" {
		log.Fatal(500, "l'article n'est pas valide")
	}
	if a.Title == nil {
		log.Fatal(500, "l'article n'a pas de titre")
	}
	a.Slug = SlugUnique(strings.ToLower(tools.SanitizerSlug(*a.Title)))
	result := db.Create(&a)

	return a, result.Error
}

func UpdateArticle(a *Article) (*Article, error) {
	if a == nil {
		log.Fatal(500, "il n'y a pas d'article a update")
	}
	a.Slug = SlugUnique(strings.ToLower(tools.SanitizerSlug(*a.Title)))
	result := db.Save(&a)

	return a, result.Error
}

func SlugUnique(title string) string {
	slugValid := false
	indexSlug := 1
	slug := strings.ToLower(tools.SanitizerSlug(title))

	for !slugValid {
		article, result := GetArticleBySlug(slug)
		if article.Slug != "" && result != nil {
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

func DeleteArticle(article *Article) (*Article, error) {
	result := db.Delete(&article)
	return article, result.Error
}
