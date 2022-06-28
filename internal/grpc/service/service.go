package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Pavlico/topcoin/internal/conf"
	protos "github.com/Pavlico/topcoin/internal/grpc/protos/coins"
	"github.com/Pavlico/topcoin/internal/utils/prettifier"
	"google.golang.org/grpc"
)

func GetCoins() ([]byte, int) {
	coins, err := doRequest()
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	result, err := prettifier.PrettyPrint(coins)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return result, http.StatusOK

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
