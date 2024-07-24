package services

import "github.com/RasoulZamani/hiGin/internal/modules/article/responses"

type ArticleServiceInterface interface {
	GetFeaturedArticles() responses.Articles
	GetUsualArticles() responses.Articles
}
