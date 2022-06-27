package handler

import (
	"context"

	protos "github.com/Pavlico/topcoin/internal/grpc/protos/coins"
	"github.com/Pavlico/topcoin/internal/grpc/service"
)

type Coins struct {
}

func NewCoin() *Coins {
	return &Coins{}
}
func (c *Coins) GetCoins(ctx context.Context, cr *protos.CoinRequest) (*protos.TopCoinsResponse, error) {
	result, err := service.GetCoins()
	if err != nil {
		return nil, err
	}
	var coins []*protos.CoinData
	for _, v := range result {
		coins = append(coins, &protos.CoinData{Symbol: v.Symbol, Rank: int32(v.Rank), Score: v.Score})
	}
	return &protos.TopCoinsResponse{Coins: coins}, nil
}
