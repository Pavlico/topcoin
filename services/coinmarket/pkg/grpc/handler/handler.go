package handler

import (
	"context"

	protos "github.com/Pavlico/topcoin/services/coinmarket/pkg/grpc/protos/coinmarket"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
)

type Coinmarket struct {
}

func NewScoreList() *Coinmarket {
	return &Coinmarket{}
}
func (c *Coinmarket) GetScore(ctx context.Context, cr *protos.ScoreRequest) (*protos.ScoreResponse, error) {
	result, err := score.GetScoreData(cr.Symbols)
	if err != nil {
		return nil, err
	}
	var scores map[string]*protos.ScoreData
	for k, v := range result {
		scores[k] = &protos.ScoreData{Symbol: v.Symbol, Score: v.Score}
	}
	return &protos.ScoreResponse{Scores: scores}, nil
}
