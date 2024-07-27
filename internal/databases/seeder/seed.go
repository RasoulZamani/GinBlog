package seeder

import (
	"log"
	"strconv"

	articleModel "github.com/RasoulZamani/hiGin/internal/modules/article/models"
	userModel "github.com/RasoulZamani/hiGin/internal/modules/user/models"
	"github.com/RasoulZamani/hiGin/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	// seeding database
	db := database.Connection()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("example_password"), 12)
	if err != nil {
		log.Fatal("Could not hash the password")
		return
	}
	exampleEmail := "example_user@example.com"
	user := userModel.User{
		UserName: "example_user",
		Email:    exampleEmail,
		Password: string(hashedPassword),
	}
	log.Printf("User with username: %s created", user.UserName)

	db.Create(&user)

	// create  some article for this example user
	for i := 0; i < 5; i++ {
		article := articleModel.Article{
			Title:   "Example Title_" + strconv.Itoa(i),
			Content: "Example content_" + strconv.Itoa(i),
			UserID:  1,
		}
		db.Create(&article)
		log.Printf("Article with title: %s created.", article.Title)
	}

	log.Println("Seeding db with one user and some article have done successfully!")
}
