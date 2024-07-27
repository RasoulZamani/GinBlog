package services

import (
	"github.com/RasoulZamani/hiGin/internal/modules/article/requests/articles"
	"github.com/RasoulZamani/hiGin/internal/modules/article/responses"
	UserResponse "github.com/RasoulZamani/hiGin/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() responses.Articles
	GetUsualArticles() responses.Articles
	Find(id int) (responses.Article, error)
	Store(request articles.StoreRequest, user UserResponse.User) (responses.Article, error)
}
