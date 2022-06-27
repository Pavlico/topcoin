package service

import (
	"context"
	"time"

	"github.com/Pavlico/topcoin/services/topcollector/pkg/conf"
	protosCoinmarket "github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/protos/coinmarket"
	protosCryptocompare "github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/protos/cryptocompare"
	protos "github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/protos/topcollector"
	"google.golang.org/grpc"
)

func GetMergedData() ([]*protos.TopCoinData, error) {
	topData, err := getTop()
	if err != nil {
		//still thinking how to log err
	}
	var symbols []string
	for symbol := range topData.Ranks {
		symbols = append(symbols, symbol)
	}
	scoreData, err := getScores(symbols)
	if err != nil {
		//still thinking how to log err
	}
	mergedData, err := assembleData(topData.Ranks, scoreData.Scores)
	if err != nil {
		//still thinking how to log err
	}
	return mergedData, nil
}

func getTop() (*protosCryptocompare.TopResponse, error) {
	conn, err := grpc.Dial(conf.CryptocompareUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := protosCryptocompare.NewCryptocompareClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetTop(ctx, &protosCryptocompare.TopRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func getScores(symbols []string) (*protosCoinmarket.ScoreResponse, error) {
	conn, err := grpc.Dial(conf.CoinmarketUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := protosCoinmarket.NewCoinmarketClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetScore(ctx, &protosCoinmarket.ScoreRequest{Symbols: symbols})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func assembleData(topData map[string]*protosCryptocompare.TopData, scoreData map[string]*protosCoinmarket.ScoreData) ([]*protos.TopCoinData, error) {
	mergedData := []*protos.TopCoinData{}
	for symbol, data := range topData {
		scoreVal, ok := scoreData[symbol]
		if !ok {
			scoreVal = &protosCoinmarket.ScoreData{}
		}
		mergedData = append(mergedData, mergeData(data.Symbol, data.Rank, scoreVal.Score))

	}
	return mergedData, nil
}

func mergeData(symbol string, rank int32, score float32) *protos.TopCoinData {
	return &protos.TopCoinData{
		Symbol: symbol,
		Rank:   rank,
		Score:  score,
	}
}
