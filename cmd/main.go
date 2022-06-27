package main

import (
	grpcserver "github.com/Pavlico/topcoin/internal/grpc/server"
	httpserver "github.com/Pavlico/topcoin/internal/http/server"
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
