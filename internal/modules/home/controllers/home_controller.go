package controllers

import (
	"net/http"

	"github.com/RasoulZamani/hiGin/config"
	ArticleRepository "github.com/RasoulZamani/hiGin/internal/modules/article/repositories"
	"github.com/gin-gonic/gin"
)

type controller struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *controller {
	return &controller{
		articleRepository: ArticleRepository.New(),
	}
}

func (controller *controller) Index(c *gin.Context) {
	// myEnv := config.ReadEnv()
	// // c.JSON(200, gin.H{
	// html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	// 	"APP_NAME": myEnv["APP_NAME"],
	// 	"message":  "Welcome Home",
	// })
	c.JSON(http.StatusOK, gin.H{
		"articles": controller.articleRepository.List(3),
	})
}

func (controller *controller) About(c *gin.Context) {
	myEnv := config.ReadEnv()
	c.JSON(200, gin.H{
		"APP_NAME": myEnv["APP_NAME"],
		"message":  "about this app",
	})
}
