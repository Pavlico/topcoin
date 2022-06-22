package grpcserver

import (
	"context"

	protos "github.com/Pavlico/topcoin/services/cryptocompare/pkg/protos/cryptocompare"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/top"
)

type Cryptocompare struct {
}

func NewTopList() *Cryptocompare {
	return &Cryptocompare{}
}
func (c *Cryptocompare) GetTop(ctx context.Context, cr *protos.TopRequest) (*protos.TopResponse, error) {
	result, err := top.GetTopData()
	if err != nil {
		return nil, err
	}
	var toplist map[string]*protos.TopData
	for k, v := range result {
		toplist[k] = &protos.TopData{Symbol: v.Symbol, Rank: int32(v.Rank)}
	}
	return &protos.TopResponse{Ranks: toplist}, nil
}
