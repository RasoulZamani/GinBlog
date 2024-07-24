package bootstrap

import (
	"github.com/RasoulZamani/hiGin/internal/databases/seeder"
	"github.com/RasoulZamani/hiGin/pkg/database"
)

func Seed() {

	database.Connect()
	seeder.Seed()

}
