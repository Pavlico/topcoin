package main

import (
	"github.com/Pavlico/topcoin/internal/server"
)

func main() {

	serverManager := server.InitServer()
	serverManager.Serve()
}
