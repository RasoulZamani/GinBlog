package routes

import (
	UserController "github.com/RasoulZamani/hiGin/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	userController := UserController.New()
	route.GET("/register", userController.Register)
	route.POST("/register", userController.HandleRegister)

}
