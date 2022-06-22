package grpcserver

import (
	"context"

	protos "github.com/Pavlico/topcoin/services/coinmarket/pkg/protos/coinmarket"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/top"
)

type Coinmarket struct {
}

func NewScoreList() *Coinmarket {
	return &Coinmarket{}
}
func (c *Coinmarket) GetScore(ctx context.Context, cr *protos.ScoreRequest) (*protos.ScoreResponse, error) {
	topData, err := top.GetTopData()
	if err != nil {
		return nil, err
	}
	var symbols []string
	for symbol := range topData {
		symbols = append(symbols, symbol)
	}
	result, err := score.GetScoreData(symbols)
	var scores map[string]*protos.ScoreData
	for k, v := range result {
		scores[k] = &protos.ScoreData{Symbol: v.Symbol, Score: v.Score}
	}
	return &protos.ScoreResponse{Scores: scores}, nil
}
