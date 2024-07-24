package controllers

import (
	"net/http"
	"strconv"

	ArticleService "github.com/RasoulZamani/hiGin/internal/modules/article/services"
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
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Could not catch article id from url and convert it to integer",
			"error":   err.Error(),
		})
		return
	}

	// find article by id
	article, err := controller.articleService.Find(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "article founded",
		"article": article,
	})
}
