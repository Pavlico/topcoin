package main

import (
	"net"
	"os"

	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/grpcserver"
	protos "github.com/Pavlico/topcoin/services/cryptocompare/pkg/protos/cryptocompare"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	cs := grpcserver.NewTopList()
	protos.RegisterCryptocompareServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", ":8061")
	if err != nil {
		os.Exit(1)
	}
	gs.Serve(l)
}
