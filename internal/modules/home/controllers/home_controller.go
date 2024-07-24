package controllers

import (
	"net/http"

	"github.com/RasoulZamani/hiGin/config"
	ArticleService "github.com/RasoulZamani/hiGin/internal/modules/article/services"
	"github.com/gin-gonic/gin"
)

type controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *controller {
	return &controller{
		articleService: ArticleService.New(),
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
		"featured_articles": controller.articleService.GetFeaturedArticles(),
		"usual_articles":    controller.articleService.GetUsualArticles(),
	})
}

func (controller *controller) About(c *gin.Context) {
	myEnv := config.ReadEnv()
	c.JSON(200, gin.H{
		"APP_NAME": myEnv["APP_NAME"],
		"message":  "about this app",
	})
}
