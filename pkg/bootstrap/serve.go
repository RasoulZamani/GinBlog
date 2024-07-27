package bootstrap

import (
	static "github.com/RasoulZamani/hiGin/pkg"
	"github.com/RasoulZamani/hiGin/pkg/database"
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/RasoulZamani/hiGin/pkg/routing"
	"github.com/RasoulZamani/hiGin/pkg/sessions"
)

func Serve() {

	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())

	html.LoadHtml(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
