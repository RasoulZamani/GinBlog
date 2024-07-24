package services

import (
	"github.com/RasoulZamani/hiGin/internal/modules/article/models"
	"github.com/RasoulZamani/hiGin/internal/modules/article/repositories"
)

type ArticleService struct {
	articleRepository repositories.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: repositories.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() []models.Article {
	return articleService.articleRepository.List(2)
}

func (articleService *ArticleService) GetUsualArticles() []models.Article {
	return articleService.articleRepository.List(3)
}
