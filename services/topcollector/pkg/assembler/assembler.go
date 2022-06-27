package assembler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Pavlico/topcoin/services/topcollector/pkg/conf"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/database"
	errorsPkg "github.com/pkg/errors"
)

func Get(outputChan chan<- []dataTypes.CoinData, errorChan chan<- error, ctx context.Context) {
	db, err := database.Initialize()
	if err != nil {
		errorChan <- err
	}
	topData, err := getTopData()
	if err != nil {
		errorChan <- err
	}
	var symbols []string
	for _, v := range topData {
		symbols = append(symbols, v.Symbol)
	}
	scoreData, err := getScoreData(symbols)
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
func getTopData() (map[string]*dataTypes.TopData, error) {
	client := getClient()
	tResponse := make(map[string]*dataTypes.TopData)
	req, err := createRequest(conf.CryptocompareUrl+conf.CryptocompareTopEndpoint, []string{})
	if err != nil {
		return tResponse, err
	}
	response, err := getResponse(req, client)
	if err != nil {
		return tResponse, err
	}
	err = json.Unmarshal(response, &tResponse)
	if err != nil {
		return tResponse, err
	}
	return tResponse, nil
}

func getScoreData(symbols []string) (map[string]*dataTypes.ScoreData, error) {
	client := getClient()
	sResponse := make(map[string]*dataTypes.ScoreData)
	req, err := createRequest(conf.CoinmarketUrl+conf.CoinmarketScoreEndpoint, symbols)
	if err != nil {
		return nil, err
	}
	response, err := getResponse(req, client)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &sResponse)
	if err != nil {
		return nil, err
	}
	return sResponse, nil
}

func AssembleData(topData map[string]*dataTypes.TopData, scoreData map[string]*dataTypes.ScoreData) ([]dataTypes.CoinData, error) {
	mergedData := []dataTypes.CoinData{}
	for _, data := range topData {
		scoreVal, ok := scoreData[data.Symbol]
		if !ok {
			scoreVal = &dataTypes.ScoreData{}
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

func createRequest(path string, symbols []string) (*http.Request, error) {
	param := url.Values{}

	if len(symbols) > 0 {
		param.Add(conf.SymbolParam, strings.Join(symbols, ","))
	}
	url := path + param.Encode()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil

}

func getResponse(req *http.Request, client http.Client) ([]byte, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "Error during sendiNG request")
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "Error during reading response data")
	}
	return responseData, nil
}

func getClient() http.Client {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	return client
}
