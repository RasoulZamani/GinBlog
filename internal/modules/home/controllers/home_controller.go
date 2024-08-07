package controllers

import (
	"net/http"

	"github.com/RasoulZamani/hiGin/config"
	ArticleService "github.com/RasoulZamani/hiGin/internal/modules/article/services"
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	myEnv := config.ReadEnv()
	// c.JSON(200, gin.H{
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"APP_NAME":          myEnv["APP_NAME"],
		"message":           "Welcome Home",
		"featured_articles": controller.articleService.GetFeaturedArticles(),
		"story_articles":    controller.articleService.GetUsualArticles(),
	})
}

func (controller *Controller) About(c *gin.Context) {
	myEnv := config.ReadEnv()
	c.JSON(200, gin.H{
		"APP_NAME": myEnv["APP_NAME"],
		"message":  "about this app",
	})
}
