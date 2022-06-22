package main

import (
	"net"
	"os"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/grpcserver"
	protos "github.com/Pavlico/topcoin/services/coinmarket/pkg/protos/coinmarket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	cs := grpcserver.NewScoreList()
	protos.RegisterCoinmarketServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", ":8061")
	if err != nil {
		os.Exit(1)
	}
	gs.Serve(l)
}
