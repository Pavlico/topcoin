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

	"github.com/Pavlico/topcoin/internal/http/handler"
)

const (
	grpcType    = "grpc"
	defaultType = "http"
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

func shutdown(ctx context.Context, server *http.Server) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		log.Println("service shutdowned")
	}
}

func (ss *ServerStarterStruct) Serve() {
	routes := http.NewServeMux()
	if ss.communicationType == defaultType {
		routes.HandleFunc("/top", handler.GetCoinsHttp())
	}
	if ss.communicationType == grpcType {
		routes.HandleFunc("/top", handler.GetCoinsGrpc())
	}
	s := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	go start(s)
	stopCh, closeCh := createChannel()
	defer shutdown(context.Background(), s)

	defer closeCh()
	log.Println("notified:", <-stopCh)
}

func CheckConnections() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
}
