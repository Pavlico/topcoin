package main

import (
	grpcserver "github.com/Pavlico/topcoin/services/coinmarket/pkg/grpc/server"
	httpserver "github.com/Pavlico/topcoin/services/coinmarket/pkg/http/server"
)

func main() {
	if true {
		serverManager := httpserver.InitServer()
		serverManager.Serve()
	}
	if false {
		grpcserver.Serve()
	}
}
