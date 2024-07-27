package middlewares

import (
	"net/http"
	"strconv"

	"github.com/RasoulZamani/hiGin/internal/modules/user/repositories"
	"github.com/RasoulZamani/hiGin/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {
	var userRepository = repositories.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepository.FindByID(userID)
		if user.ID != 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.Next()

	}
}
