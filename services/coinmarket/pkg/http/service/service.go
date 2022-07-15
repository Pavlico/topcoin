package service

import (
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
)

func GetScores(symbols []string) (map[string]dataTypes.ScoreData, error) {
	topData, err := score.GetScoreData(symbols)
	if err != nil {
		return nil, err
	}

	return topData, nil
}
