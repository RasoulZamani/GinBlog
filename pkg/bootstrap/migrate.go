package bootstrap

import (
	"github.com/RasoulZamani/hiGin/internal/databases/migration"
	"github.com/RasoulZamani/hiGin/pkg/database"
)

func Migrate() {

	database.Connect()
	migration.Migrate()

}
