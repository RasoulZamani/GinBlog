package responses

import "github.com/RasoulZamani/hiGin/internal/modules/user/models"

type User struct {
	ID       uint
	Username string
	Email    string
}

type Users struct {
	Data []User
}

func ToUser(user models.User) User {
	return User{
		ID:       user.ID,
		Username: user.UserName,
		Email:    *user.Email,
	}
}
