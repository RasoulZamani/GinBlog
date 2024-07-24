package routes

import (
	ArticleController "github.com/RasoulZamani/hiGin/internal/modules/article/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	articleController := ArticleController.New()
	route.GET("/articles/:id", articleController.Show)
}
