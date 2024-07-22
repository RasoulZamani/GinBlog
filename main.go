package main

import (
	"fmt"

	"github.com/RasoulZamani/hiGin/config"
	"github.com/gin-gonic/gin"
)

func main() {
	myEnv := config.ReadEnv()
	fmt.Println(myEnv)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app name": myEnv["APP_NAME"],
			"message":  "pong",
		})
	})
	r.Run(myEnv["HOST"] + ":" + myEnv["PORT"]) // listen and serve on 0.0.0.0:8080
}
