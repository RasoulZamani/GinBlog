package controllers

import (
	"net/http"
	"strconv"

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

func (controller *Controller) Show(c *gin.Context) {

	// catch id and convert to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// c.JSON(http.StatusUnprocessableEntity, gin.H{
		// 	"message": "Could not catch article id from url and convert it to integer",
		// 	"error":   err.Error(),
		// })
		html.Render(
			c,
			http.StatusUnprocessableEntity,
			"templates/errors/html/422",
			gin.H{
				"title":   "bad request",
				"message": "Could not catch article id from url and convert it to integer",
			},
		)
		return
	}

	// find article by id
	article, err := controller.articleService.Find(id)
	if err != nil {
		// c.JSON(http.StatusNotFound, gin.H{
		// 	"message": err.Error(),
		// })
		html.Render( // show 404 page if not found
			c,
			http.StatusNotFound,
			"templates/errors/html/404",
			gin.H{
				"title":   "Page not found",
				"message": err.Error(),
			},
		)
		return

	}

	html.Render(c, http.StatusNotFound, "modules/article/html/show", gin.H{
		"title":   "Show article",
		"article": article,
	})
	return
}

// ********************** Create Article ***********************
func (controller *Controller) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "article creation form",
	})
}

func (controller *Controller) Store(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "article created",
	})
}
