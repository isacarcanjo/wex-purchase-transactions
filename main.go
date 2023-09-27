package main

import (
	"fmt"
	_ "serve/docs"
	. "serve/v1"
	. "serve/v1/config"
	"time"
)

// @title Purchase Transaction API Docs
// @version 1.0.0
// @contact.name Isac Arcanjo
// @contact.url    https://github.com/isacarcanjo

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000
// @BasePath  /v1
func main() {
	config := GetDotEnv()

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	fmt.Println("Brazil time:", time.Now())

	severHttp := HTTP{Config: config}
	startServer(severHttp)
}

func startServer(s Server) {
	s.Start()
}
