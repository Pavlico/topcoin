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

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/conf"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/http/handler"
)

func createChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return stopCh, func() {
		close(stopCh)
	}
}

func start(server *http.Server) error {
	log.Println("service started")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	} else {
		log.Println("service stopped gracefully")
		return nil
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
	routes := http.NewServeMux()
	routes.HandleFunc(conf.ServiceConfig.HttpEndpoint, handler.GetScores())
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
