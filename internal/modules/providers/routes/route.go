package routes

import (
	HomeRoutes "github.com/RasoulZamani/hiGin/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	HomeRoutes.Routes(router)

}
