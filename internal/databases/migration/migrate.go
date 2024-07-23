package migration

import (
	"log"

	articleModels "github.com/RasoulZamani/hiGin/internal/modules/article/models"
	userModels "github.com/RasoulZamani/hiGin/internal/modules/user/models"
	"github.com/RasoulZamani/hiGin/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})
	if err != nil {
		log.Fatal("Could not migrate models")
		return
	}

	log.Println("Migrate models successfully!")
}
