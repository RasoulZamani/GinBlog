package bootstrap

import (
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/RasoulZamani/hiGin/pkg/routing"
)

func Serve() {
	routing.Init()
	routing.RegisterRoutes()
	html.LoadHtml(routing.GetRouter())
	routing.Serve()
}
