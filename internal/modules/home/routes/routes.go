package routes

import (
	"time"

	"github.com/RasoulZamani/hiGin/config"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	myEnv := config.ReadEnv()

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time":     time.Now(),
			"app name": myEnv["APP_NAME"],
			"message":  "Welcome Home",
		})
	})
}
