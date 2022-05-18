package assembler

import (
	"context"
	"topcoin/internal/dataTypes"
	"topcoin/internal/score"
	"topcoin/internal/top"
)

func Get(reqTypes []string, outputChan chan<- []dataTypes.CoinData, errorChan chan<- error, ctx context.Context) {
	topData, err := top.Process()
	if err != nil {
		errorChan <- err
	}
	scoreData, err := score.Process(topData)
	if err != nil {
		//place for logger
	}
	mergedData, err := AssembleData(topData, scoreData)
	if err != nil {
		errorChan <- err
	}
	outputChan <- mergedData

}

func AssembleData(topData map[string]dataTypes.TopData, scoreData map[string]dataTypes.ScoreData) ([]dataTypes.CoinData, error) {
	mergedData := []dataTypes.CoinData{}
	for symbol, data := range topData {
		scoreVal, ok := scoreData[symbol]
		if !ok {
			scoreVal = dataTypes.ScoreData{}
		}
		mergedData = append(mergedData, mergeData(data.Symbol, data.Rank, scoreVal.Score))

	}
	return mergedData, nil
}

func mergeData(symbol string, rank int, score float32) dataTypes.CoinData {
	return dataTypes.CoinData{
		Symbol: symbol,
		Rank:   rank,
		Score:  score,
	}
}
