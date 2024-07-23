package bootstrap

import (
	static "github.com/RasoulZamani/hiGin/pkg"
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/RasoulZamani/hiGin/pkg/routing"
)

func Serve() {
	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHtml(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}