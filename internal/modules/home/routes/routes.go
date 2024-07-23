package routes

import (
	HomeController "github.com/RasoulZamani/hiGin/internal/modules/home/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	homeController := HomeController.New()
	route.GET("/", homeController.Index)
	route.GET("/about", homeController.About)
}
