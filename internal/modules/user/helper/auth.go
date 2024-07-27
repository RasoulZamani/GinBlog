package helper

import (
	"strconv"

	UserRepository "github.com/RasoulZamani/hiGin/internal/modules/user/repositories"
	"github.com/RasoulZamani/hiGin/internal/modules/user/responses"
	"github.com/RasoulZamani/hiGin/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) responses.User {
	var response responses.User

	// read user id of authenticated user from sessions
	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)

	// find all data of user from db by id
	userRepository := UserRepository.New()
	user := userRepository.FindByID(userID)

	if user.ID == 0 {
		return response // empty response
	}
	return responses.ToUser(user)
}
