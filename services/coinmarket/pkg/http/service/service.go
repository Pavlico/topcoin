package service

import (
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/utils/prettifier"
)

func GetScores(symbols []string) ([]byte, error) {
	topData, err := score.GetScoreData(symbols)
	if err != nil {
		return nil, err
	}
	result, err := prettifier.PrettyPrint(topData)
	if err != nil {
		return nil, err
	}
	return result, nil
}
