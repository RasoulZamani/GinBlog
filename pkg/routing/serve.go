package routing

import (
	"log"

	"github.com/RasoulZamani/hiGin/config"
)

func Serve() {
	r := GetRouter()
	myEnv := config.ReadEnv()

	err := r.Run(myEnv["HOST"] + ":" + myEnv["PORT"]) // listen and serve on host:port

	if err != nil {
		log.Fatal("Can not run server for this routing")
	} else {
		log.Println("server running")
	}
}
