package main

import (
	"log"

	httpserver "github.com/Pavlico/topcoin/internal/http/server"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/flags"
)

func main() {
	flagData := flags.GetFlagData()
	if err := flagData.ValidateFlags(); err != nil {
		log.Println(err)
		return
	}
	serverManager := httpserver.InitServer(flagData.CommunicationType)
	serverManager.Serve()

}
