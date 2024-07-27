package repositories

import (
	"log"

	"github.com/RasoulZamani/hiGin/internal/modules/user/models"
	"github.com/RasoulZamani/hiGin/pkg/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// new instance of repository
func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

// method for creating user
func (userRepository *UserRepository) Create(user models.User) models.User {
	var newUser models.User
	userRepository.DB.Create(&user).Scan(&newUser)
	return newUser
}

func (userRepository *UserRepository) FindByEmail(email string) models.User {
	var user models.User
	userRepository.DB.First(&user, "email = ?", email)
	log.Println("Id of user is", user.ID)
	return user
}
