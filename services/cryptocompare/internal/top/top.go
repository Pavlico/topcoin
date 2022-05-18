package top

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Pavlico/topcoin/services/cryptocompare/internal/conf"
	"github.com/Pavlico/topcoin/services/cryptocompare/internal/dataTypes"
)

type TopResponse struct {
	Data []TopResponseData `json:"Data"`
	TopResponseError
}
type TopResponseData struct {
	CoinInfo struct {
		Symbol string `json:"Name"`
	} `json:"CoinInfo"`
}
type TopResponseError struct {
	Message string `json:"Message"`
}

func Process() (map[string]dataTypes.TopData, error) {
	client := getClient()
	tResponse := TopResponse{}
	topData := make(map[string]dataTypes.TopData)
	apiConf := conf.ApiConfig[conf.TopApi]
	pageNumInt, err := strconv.Atoi(apiConf.Options[conf.PageParam])
	if err != nil {
		return nil, err
	}
	for i := 0; i < pageNumInt; i++ {
		currentPage := strconv.Itoa(i)
		req, err := tResponse.CreateRequest(apiConf, currentPage)
		if err != nil {
			return nil, err
		}
		response, err := tResponse.GetResponse(req, client)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(response, &tResponse)
		if err != nil {
			return nil, err
		}
		err = tResponse.ValidateResponse()
		if err != nil {
			return nil, err
		}
		for _, v := range tResponse.Data {
			topData[v.CoinInfo.Symbol] = dataTypes.TopData{Symbol: v.CoinInfo.Symbol, Rank: len(topData) + 1}
		}
	}
	return topData, nil
}

func (tResponse TopResponse) ValidateResponse() error {
	if tResponse.TopResponseError.Message != conf.SuccessMessage {
		return errors.New("Response TOP is not valid")
	}
	return nil
}

func (tResponse TopResponse) GetResponse(req *http.Request, client http.Client) ([]byte, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return responseData, nil
}

func (tr TopResponse) CreateRequest(apiData conf.ApiData, currentPage string) (*http.Request, error) {
	param := url.Values{}
	apiData.Options[conf.PageParam] = currentPage
	for i, val := range apiData.Options {
		param.Add(i, val)
	}

	endpoint := apiData.ApiAddress + apiData.EndPoint + param.Encode()
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(apiData.CredentialsHeader, apiData.Credentials)
	return req, nil
}

func getClient() http.Client {
	client := http.Client{
		Timeout: conf.ApiTimeout,
	}
	return client
}
