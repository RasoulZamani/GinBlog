package repositories

import "github.com/RasoulZamani/hiGin/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []models.Article
	Find(id int) models.Article
}
