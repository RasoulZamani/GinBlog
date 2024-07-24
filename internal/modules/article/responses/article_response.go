package responses

import (
	"fmt"
	"strconv"

	"math/rand/v2"

	ArticleModel "github.com/RasoulZamani/hiGin/internal/modules/article/models"
	UserResponse "github.com/RasoulZamani/hiGin/internal/modules/user/responses"
)

type Article struct {
	ID        uint
	Title     string
	Image     string
	Content   string
	CreatedAt string
	User      UserResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article ArticleModel.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Image:     "/assets/img/demopic/" + strconv.Itoa(rand.IntN(10)) + ".jpg",
		Content:   article.Content,
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      UserResponse.ToUser(article.User),
	}
}

func ToArticles(articles []ArticleModel.Article) Articles {
	var response Articles
	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}
	return response
}
