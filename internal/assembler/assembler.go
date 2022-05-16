package assembler

import (
	"context"
	"encoding/json"
	"time"
	"topcoin/internal/score"
	"topcoin/internal/top"
)

type CoinData struct {
	Symbol string  `json:"Symbol"`
	Rank   int     `json:"Rank"`
	Score  float32 `json:"Score"`
}

func Get(reqTypes []string, outputChan chan<- string, errorChan chan<- error, ctx context.Context) {
	topRequestChan := make(chan map[string]top.TopData)
	scoreRequestChan := make(chan map[string]score.ScoreData)
	var topData map[string]top.TopData
	var scoreData map[string]score.ScoreData
	for {
		go top.Process(topRequestChan, errorChan)
		go score.Process(scoreRequestChan, errorChan)
		select {
		case topData = <-topRequestChan:
		case scoreData = <-scoreRequestChan:
		case <-time.After(1 * time.Second):
		case <-ctx.Done():
			return
		}
		// why not like that? :)
		// topData = <-topRequestChan:
		// scoreData = <-scoreRequestChan:
		if len(topData) == 0 || len(topData) == 0 {
			continue
		}
		mergedData, err := AssembleData(topData, scoreData)
		if err != nil {
			errorChan <- err
		}
		prettyMergedMap, err := PrettyPrint(mergedData)
		if err != nil {
			errorChan <- err
		}
		outputChan <- string(prettyMergedMap)

	}

}

func AssembleData(topData map[string]top.TopData, scoreData map[string]score.ScoreData) ([]CoinData, error) {
	mergedData := []CoinData{}
	for symbol, data := range topData {
		scoreVal, ok := scoreData[symbol]
		if !ok {
			scoreVal = score.ScoreData{}
		}
		mergedData = append(mergedData, mergeData(data.Symbol, data.Rank, scoreVal.Score))

	}
	return mergedData, nil
}

func PrettyPrint(i interface{}) ([]byte, error) {
	return json.MarshalIndent(i, "", "\t")
}

func mergeData(symbol string, rank int, score float32) CoinData {
	return CoinData{
		Symbol: symbol,
		Rank:   rank,
		Score:  score,
	}
}
