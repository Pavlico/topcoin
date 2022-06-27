package server

import (
	"net"
	"os"

	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/conf"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/grpc/handler"
	protos "github.com/Pavlico/topcoin/services/cryptocompare/pkg/grpc/protos/cryptocompare"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve() {
	gs := grpc.NewServer()
	cs := handler.NewTopList()
	protos.RegisterCryptocompareServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", conf.TopPort)
	if err != nil {
		os.Exit(1)
	}
	gs.Serve(l)
}
