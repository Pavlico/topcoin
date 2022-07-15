package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pavlico/topcoin/internal/conf"
	"github.com/Pavlico/topcoin/internal/http/handler"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/flags"
)

const (
	grpcType    = "grpc"
	httpType    = "http"
	defaultType = httpType
)

type ServerStarter interface {
	Serve(server *http.Server)
}

type ServerStarterStruct struct {
	communicationType string
}

func InitServer(communicationType string) *ServerStarterStruct {
	return &ServerStarterStruct{
		communicationType: communicationType,
	}
}

func createChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return stopCh, func() {
		close(stopCh)
	}
}

func start(server *http.Server) {
	log.Println("service started")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	} else {
		log.Println("service stopped gracefully")
	}
}

func shutdown(ctx context.Context, server *http.Server) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return err
	} else {
		log.Println("service shutdowned")
		return nil
	}
}

func Serve() error {
	flagData := flags.GetFlagData()
	if err := flagData.ValidateFlags(); err != nil {
		return err
	}
	routes := http.NewServeMux()
	if flagData.CommunicationType == defaultType {
		routes.HandleFunc(conf.ServiceConfig.HttpEndpoint, handler.GetCoinsHttp())
	}
	if flagData.CommunicationType == grpcType {
		routes.HandleFunc(conf.ServiceConfig.HttpEndpoint, handler.GetCoinsGrpc())
	}
	s := &http.Server{
		Addr:    conf.ServiceConfig.HttpPort,
		Handler: routes,
	}
	go start(s)
	stopCh, closeCh := createChannel()
	<-stopCh
	closeCh()
	err := shutdown(context.Background(), s)
	if err != nil {
		return err
	}
	return nil
}

func CheckConnections() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
}
