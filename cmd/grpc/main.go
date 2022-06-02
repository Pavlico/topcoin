package main

import (
	"net"
	"os"

	"github.com/Pavlico/topcoin/internal/grpcserver"
	protos "github.com/Pavlico/topcoin/internal/protos/coins"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	cs := grpcserver.NewCoin()
	protos.RegisterCoinsServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		os.Exit(1)
	}
	gs.Serve(l)
}
