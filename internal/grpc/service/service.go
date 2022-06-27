package service

import (
	"context"
	"time"

	"github.com/Pavlico/topcoin/internal/conf"
	protos "github.com/Pavlico/topcoin/internal/grpc/protos/coins"
	"google.golang.org/grpc"
)

func GetCoins() ([]*protos.CoinData, error) {
	coins, err := doRequest()
	if err != nil {
		return nil, err
	}
	return coins.Coins, nil

}

func doRequest() (*protos.TopCoinsResponse, error) {
	conn, err := grpc.Dial(conf.TopCollectorEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := protos.NewCoinsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetCoins(ctx, &protos.CoinRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}
