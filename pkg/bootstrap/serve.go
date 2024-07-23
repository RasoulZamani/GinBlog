package bootstrap

import (
	"github.com/RasoulZamani/hiGin/pkg/routing"
)

func Serve() {
	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
