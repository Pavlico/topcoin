package server

import (
	"net"
	"os"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/grpc/handler"
	protos "github.com/Pavlico/topcoin/services/coinmarket/pkg/grpc/protos/coinmarket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve() {
	gs := grpc.NewServer()
	cs := handler.NewScoreList()
	protos.RegisterCoinmarketServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", ":8061")
	if err != nil {
		os.Exit(1)
	}
	gs.Serve(l)
}
