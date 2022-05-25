package assembler

import (
	"errors"

	scoreTypes "github.com/Pavlico/topcoin/services/coinmarket/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
	topTypes "github.com/Pavlico/topcoin/services/cryptocompare/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/top"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/conf"
	coinTypes "github.com/Pavlico/topcoin/services/topcollector/pkg/dataTypes"
)

func Get(reqTypes []string) (map[int]*coinTypes.CoinData, error) {
	dataToMerge := make(map[string]interface{})
	var topData map[string]topTypes.TopData
	for _, v := range reqTypes {
		switch v {
		case conf.TopApi:
			topData, err := top.GetTopData()
			if err != nil {
				return nil, err
			}
			dataToMerge[v] = topData
		case conf.ScoreApi:
			_, ok := dataToMerge["top"]
			if !ok {
				return nil, errors.New("No top data")
			}
			var symbols []string
			for symbol, _ := range topData {
				symbols = append(symbols, symbol)
			}
			scoreData, err := score.GetScoreData(symbols)
			if err != nil {
				return nil, err
			}
			dataToMerge[v] = scoreData
		}
	}

	mergedData, err := AssembleData(dataToMerge)
	if err != nil {
		return nil, err
	}

	return mergedData, nil
}

func AssembleData(dataToMerge map[string]interface{}) (map[int]*coinTypes.CoinData, error) {
	minMergeAmount := 2
	if len(dataToMerge) < minMergeAmount {
		return nil, errors.New("Not enough data to merge")
	}
	mergedMap := make(map[int]*coinTypes.CoinData)
	nameToRankMap := make(map[string]int)
	for _, m := range dataToMerge {
		switch v := m.(type) {
		case []topTypes.TopData:
			{
				for _, k := range v {
					mergedMap[k.Rank] = &coinTypes.CoinData{
						Rank:   k.Rank,
						Symbol: k.Symbol,
					}
					nameToRankMap[k.Symbol] = k.Rank
				}
			}
		case []scoreTypes.ScoreData:
			{
				for _, k := range v {
					if rank, ok := nameToRankMap[k.Symbol]; ok {
						mergedMap[rank].Score = k.Score
					}
				}
			}
		}
	}
	return mergedMap, nil
}
