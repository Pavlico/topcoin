package main

import (
	"log"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/flags"
	grpcserver "github.com/Pavlico/topcoin/services/coinmarket/pkg/grpc/server"
	httpserver "github.com/Pavlico/topcoin/services/coinmarket/pkg/http/server"
)

const (
	grpcType    = "grpc"
	defaultType = "http"
)

func main() {
	flagData := flags.GetFlagData()
	if err := flagData.ValidateFlags(); err != nil {
		log.Println(err)
		return
	}
	if flagData.CommunicationType == defaultType {
		serverManager := httpserver.InitServer()
		serverManager.Serve()
	}
	if flagData.CommunicationType == grpcType {
		grpcserver.Serve()
	}
}
