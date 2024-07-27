package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// function for create a session as a middleware for given router;
// It will be used in bootstrap
func Start(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret_key"))
	router.Use(sessions.Sessions("sessions", store))
}
