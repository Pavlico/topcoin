package server

import (
	"net"
	"os"

	"github.com/Pavlico/topcoin/services/topcollector/pkg/conf"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/handler"
	protos "github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/protos/topcollector"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve() {
	gs := grpc.NewServer()
	cs := handler.NewCoinList()
	protos.RegisterTopCollectorServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", conf.TopCoinPort)
	if err != nil {
		os.Exit(1)
	}
	gs.Serve(l)
}
