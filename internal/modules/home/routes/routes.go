package routes

import (
	"net/http"

	"github.com/RasoulZamani/hiGin/config"
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	myEnv := config.ReadEnv()

	route.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title":   myEnv["APP_NAME"],
			"message": "Welcome Home",
		})
	})

	route.GET("/about", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title":   myEnv["APP_NAME"],
			"message": "about this app",
		})

	})
}
