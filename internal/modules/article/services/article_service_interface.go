package services

import "github.com/RasoulZamani/hiGin/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []models.Article
	GetUsualArticles() []models.Article
}
