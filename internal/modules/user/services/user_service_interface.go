package services

import (
	"github.com/RasoulZamani/hiGin/internal/modules/user/requests/auth"
	"github.com/RasoulZamani/hiGin/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (responses.User, error)
	CheckEmailExist(email string) bool
	LoginUser(request auth.LoginRequest) (responses.User, error)
}
