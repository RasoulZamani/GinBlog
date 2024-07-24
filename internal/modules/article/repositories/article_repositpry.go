package repositories

import (
	"github.com/RasoulZamani/hiGin/internal/modules/article/models"
	"github.com/RasoulZamani/hiGin/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

// new instance of repository
func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

// method for list articles
func (articleRepository *ArticleRepository) List(limit int) []models.Article {
	var articles []models.Article

	articleRepository.DB.Limit(limit).Joins("User").Find(&articles)
	return articles
}

// method for get one article by id
func (articleRepository *ArticleRepository) Find(id int) models.Article {
	var article models.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}
