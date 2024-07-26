package services

import (
	"github.com/RasoulZamani/hiGin/internal/modules/user/requests/auth"
	"github.com/RasoulZamani/hiGin/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (responses.User, error)
}
