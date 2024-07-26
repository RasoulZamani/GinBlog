package services

import (
	"errors"
	"log"

	"github.com/RasoulZamani/hiGin/internal/modules/user/models"
	"github.com/RasoulZamani/hiGin/internal/modules/user/repositories"
	"github.com/RasoulZamani/hiGin/internal/modules/user/requests/auth"
	"github.com/RasoulZamani/hiGin/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: repositories.New(),
	}
}

func (UserService *UserService) Create(request auth.RegisterRequest) (responses.User, error) {
	var response responses.User
	var user models.User

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		log.Fatal("Could not hash the password")
		return response, err
	}
	user.UserName = request.Name
	user.Email = &request.Email
	user.Password = string(hashedPassword)

	newUser := UserService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating user")
	}

	return responses.ToUser(newUser), nil
}
