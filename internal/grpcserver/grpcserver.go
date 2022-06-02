package grpcserver

import (
	"context"

	"github.com/Pavlico/topcoin/internal/coinservice"
	protos "github.com/Pavlico/topcoin/internal/protos/coins"
)

type Coins struct {
}

func NewCoin() *Coins {
	return &Coins{}
}
func (c *Coins) GetCoins(ctx context.Context, cr *protos.CoinRequest) (*protos.TopCoinsResponse, error) {
	result, err := coinservice.GetTopCoin()
	if err != nil {
		return nil, err
	}
	var coins []*protos.CoinData
	for _, v := range result {
		coins = append(coins, &protos.CoinData{Symbol: v.Symbol, Rank: int32(v.Rank), Score: v.Score})
	}
	return &protos.TopCoinsResponse{Coins: coins}, nil
}
