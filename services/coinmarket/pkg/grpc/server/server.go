package server

import (
	"net"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/conf"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/grpc/handler"
	protos "github.com/Pavlico/topcoin/services/coinmarket/pkg/grpc/protos/coinmarket"
	"google.golang.org/grpc"
)

func Serve() error {
	gs := grpc.NewServer()
	cs := handler.NewScoreList()
	protos.RegisterCoinmarketServer(gs, cs)

	l, err := net.Listen(conf.ServiceConfig.GrpcNetwork, conf.ServiceConfig.GrpcPort)
	if err != nil {
		return err
	}
	if err := gs.Serve(l); err != nil {
		return err
	}
	return nil
}
