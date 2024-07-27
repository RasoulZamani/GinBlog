package routes

import (
	"github.com/RasoulZamani/hiGin/internal/middlewares"
	ArticleController "github.com/RasoulZamani/hiGin/internal/modules/article/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articleController := ArticleController.New()
	router.GET("/articles/:id", articleController.Show)

	authGroup := router.Group("/articles")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", articleController.Create)
		authGroup.GET("/store", articleController.Store)
	}

}
