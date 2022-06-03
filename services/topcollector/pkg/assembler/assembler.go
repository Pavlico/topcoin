package assembler

import (
	"context"

	scoreTypes "github.com/Pavlico/topcoin/services/coinmarket/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
	topTypes "github.com/Pavlico/topcoin/services/cryptocompare/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/top"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/database"
)

func Get(outputChan chan<- []dataTypes.CoinData, errorChan chan<- error, ctx context.Context) {
	db, err := database.Initialize()
	if err != nil {
		errorChan <- err
	}
	topData, err := top.GetTopData()
	if err != nil {
		errorChan <- err
	}
	var symbols []string
	for symbol := range topData {
		symbols = append(symbols, symbol)
	}
	scoreData, err := score.GetScoreData(symbols)
	if err != nil {
		//place for logger
	}
	mergedData, err := AssembleData(topData, scoreData)
	if err != nil {
		errorChan <- err
	}
	for _, v := range mergedData {
		err := db.Save(v.Symbol, v.Rank, v.Score)
		if err != nil {
			errorChan <- err
		}
	}

	outputChan <- mergedData

}

func AssembleData(topData map[string]topTypes.TopData, scoreData map[string]scoreTypes.ScoreData) ([]dataTypes.CoinData, error) {
	mergedData := []dataTypes.CoinData{}
	for symbol, data := range topData {
		scoreVal, ok := scoreData[symbol]
		if !ok {
			scoreVal = scoreTypes.ScoreData{}
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
