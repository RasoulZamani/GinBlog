package services

import (
	"errors"
	"strconv"

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

func (articleService *ArticleService) Find(id int) (responses.Article, error) {
	var response responses.Article

	article := articleService.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("Article not found with this id: " + strconv.Itoa(id))
	}

	return responses.ToArticle(article), nil
}
