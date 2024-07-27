package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RasoulZamani/hiGin/internal/modules/article/requests/articles"
	ArticleService "github.com/RasoulZamani/hiGin/internal/modules/article/services"
	"github.com/RasoulZamani/hiGin/internal/modules/user/helper"
	"github.com/RasoulZamani/hiGin/pkg/converters"
	"github.com/RasoulZamani/hiGin/pkg/errors"
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/RasoulZamani/hiGin/pkg/old"
	"github.com/RasoulZamani/hiGin/pkg/sessions"
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
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create article",
	})
}

func (controller *Controller) Store(c *gin.Context) {

	// validate request
	var request articles.StoreRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)

		// save errors in session
		sessions.Set(c, "articleFormErrors", converters.MapToString(errors.Get()))
		// session then send to template by WithGlobalData function

		// save form data to session in order to re-present them if errors happened
		old.Init()
		old.Set(c)
		sessions.Set(c, "oldArticleForm", converters.ListMapToString(old.Get()))
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	// create the article
	user := helper.Auth(c)
	article, err := controller.articleService.Store(request, user)
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	log.Printf("article created successfully with id: %d", article.ID)
	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))

}
