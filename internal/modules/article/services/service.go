package services

import (
	"github.com/RasoulZamani/hiGin/internal/modules/article/repositories"
	"github.com/RasoulZamani/hiGin/internal/modules/article/responses"
)

type ArticleService struct {
	articleRepository repositories.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: repositories.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() responses.Articles {
	articles := articleService.articleRepository.List(2)

	return responses.ToArticles(articles)
}

func (articleService *ArticleService) GetUsualArticles() responses.Articles {
	articles := articleService.articleRepository.List(3)
	return responses.ToArticles(articles)
}
