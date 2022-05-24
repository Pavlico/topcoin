package assembler

import (
	"encoding/json"
	"errors"

	scoreTypes "github.com/Pavlico/topcoin/services/coinmarket/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
	topTypes "github.com/Pavlico/topcoin/services/cryptocompare/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/top"
	"github.com/Pavlico/topcoin/services/topcollector/internal/conf"
)

type PrettyResp struct {
	Symbol string  `json:"Symbol"`
	Rank   int     `json:"Rank"`
	Score  float32 `json:"Score"`
}

func Get(reqTypes []string) (interface{}, error) {
	dataToMerge := make(map[string]interface{})
	for _, v := range reqTypes {
		switch v {
		case conf.TopApi:
			data, err := top.Process()
			if err != nil {
				return nil, err
			}
			dataToMerge[v] = data
		case conf.ScoreApi:
			data, err := score.Process()
			if err != nil {
				return nil, err
			}
			dataToMerge[v] = data
		}
	}

	mergedData, err := AssembleData(dataToMerge)
	if err != nil {
		return nil, err
	}
	prettyMergedMap, err := PrettyPrint(mergedData)
	if err != nil {
		return nil, err
	}
	return string(prettyMergedMap), nil
}

func AssembleData(dataToMerge map[string]interface{}) (map[int]*PrettyResp, error) {
	minMergeAmount := 2
	if len(dataToMerge) < minMergeAmount {
		return nil, errors.New("Not enough data to merge")
	}
	mergedMap := make(map[int]*PrettyResp)
	nameToRankMap := make(map[string]int)
	for _, m := range dataToMerge {
		switch v := m.(type) {
		case []scoreTypes.ScoreData:
			{
				for _, k := range v {
					mergedMap[k.Rank] = &PrettyResp{
						Rank:   k.Rank,
						Symbol: k.Symbol,
					}
					nameToRankMap[k.Symbol] = k.Rank
				}
			}
		case []topTypes.TopData:
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

func PrettyPrint(i interface{}) ([]byte, error) {
	return json.MarshalIndent(i, "", "\t")
}
