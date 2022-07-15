package server

import (
	"net"

	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/conf"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/grpc/handler"
	protos "github.com/Pavlico/topcoin/services/cryptocompare/pkg/grpc/protos/cryptocompare"
	"google.golang.org/grpc"
)

func Serve() error {
	gs := grpc.NewServer()
	cs := handler.NewTopList()
	protos.RegisterCryptocompareServer(gs, cs)

	l, err := net.Listen(conf.ServiceConfig.GrpcNetwork, conf.ServiceConfig.GrpcPort)
	if err != nil {
		return err
	}
	if err := gs.Serve(l); err != nil {
		return err
	}
	return nil
}
