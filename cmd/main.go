package main

import (
	"github.com/Pavlico/topcoin/internal/routes"
	"github.com/Pavlico/topcoin/internal/server"
)

func main() {
	serverManager := server.ServerStarterStruct{}
	router := routes.GetAvailableRoutes()
	serverManager.Serve(router)
}
