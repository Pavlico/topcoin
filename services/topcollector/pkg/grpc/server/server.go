package server

import (
	"net"

	"github.com/Pavlico/topcoin/services/topcollector/pkg/conf"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/handler"
	protos "github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/protos/topcollector"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve() error {
	gs := grpc.NewServer()
	cs := handler.NewCoinList()
	protos.RegisterTopCollectorServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen(conf.ServiceConfig.GrpcNetwork, conf.ServiceConfig.GrpcPort)
	if err != nil {
		return err
	}
	if err := gs.Serve(l); err != nil {
		return err
	}
	return nil
}
