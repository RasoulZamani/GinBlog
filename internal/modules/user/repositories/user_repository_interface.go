package repositories

import "github.com/RasoulZamani/hiGin/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user models.User) models.User
	FindByEmail(email string) models.User
}
