package cmd

import (
	"fmt"

	"github.com/RasoulZamani/hiGin/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server and serve web app",
	Long:  `This command used to run server on host and port that declared in the .env file`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
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
