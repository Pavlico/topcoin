package top

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/conf"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/dataTypes"
)

func GetTopData() (map[string]dataTypes.TopData, error) {
	client := getClient()
	tResponse := dataTypes.TopResponse{}
	topData := make(map[string]dataTypes.TopData)
	pageNumInt, err := strconv.Atoi(conf.ServiceConfig.PageParam)
	if err != nil {
		return nil, err
	}
	for i := 0; i < pageNumInt; i++ {
		currentPage := strconv.Itoa(i)
		req, err := CreateRequest(conf.ServiceConfig, currentPage)
		if err != nil {
			return nil, err
		}
		response, err := GetResponse(req, client)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(response, &tResponse)
		if err != nil {
			return nil, err
		}
		err = ValidateResponse(tResponse)
		if err != nil {
			return nil, err
		}
		for _, v := range tResponse.Data {
			topData[v.CoinInfo.Symbol] = dataTypes.TopData{Symbol: v.CoinInfo.Symbol, Rank: len(topData) + 1}
		}
	}
	return topData, nil
}

func ValidateResponse(tResponse dataTypes.TopResponse) error {
	if tResponse.TopResponseError.Message != conf.ServiceConfig.SuccessMessage {
		return errors.New("Response TOP is not valid")
	}
	return nil
}

func GetResponse(req *http.Request, client http.Client) ([]byte, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func CreateRequest(appConfig conf.Config, currentPage string) (*http.Request, error) {
	param := url.Values{}
	param.Add(appConfig.PageParam, currentPage)
	param.Add(appConfig.TsymParam, appConfig.RankRequestTsymVal)
	param.Add(appConfig.LimitParam, appConfig.RankRequestLimitVal)

	endpoint := appConfig.RankRequestApiAddress + appConfig.RankRequestEndPoint + param.Encode()
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(appConfig.RankRequestCredentialsHeader, appConfig.RankRequestCredentials)
	return req, nil
}

func getClient() http.Client {
	client := http.Client{
		Timeout: time.Duration(conf.ServiceConfig.ApiTimeout) * time.Second,
	}
	return client
}
