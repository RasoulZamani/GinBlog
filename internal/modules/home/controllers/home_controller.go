package controllers

import (
	"net/http"

	"github.com/RasoulZamani/hiGin/config"
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/gin-gonic/gin"
)

type controller struct{}

func New() *controller {
	return &controller{}
}

func (controller *controller) Index(c *gin.Context) {
	myEnv := config.ReadEnv()
	// c.JSON(200, gin.H{
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"APP_NAME": myEnv["APP_NAME"],
		"message":  "Welcome Home",
	})
}

func (controller *controller) About(c *gin.Context) {
	myEnv := config.ReadEnv()
	c.JSON(200, gin.H{
		"APP_NAME": myEnv["APP_NAME"],
		"message":  "about this app",
	})
}
