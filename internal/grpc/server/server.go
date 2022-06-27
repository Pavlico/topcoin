package server

import (
	"net"
	"os"

	"github.com/Pavlico/topcoin/internal/grpc/handler"
	protos "github.com/Pavlico/topcoin/internal/grpc/protos/coins"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve() {
	gs := grpc.NewServer()
	cs := handler.NewCoin()
	protos.RegisterCoinsServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		os.Exit(1)
	}
	gs.Serve(l)
}
