package routing

import (
	"github.com/RasoulZamani/hiGin/internal/modules/providers/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())

}
