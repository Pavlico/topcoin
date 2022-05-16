package top

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"topcoin/internal/conf"

	errorsPkg "github.com/pkg/errors"
)

type TopData struct {
	Symbol string
	Rank   int
}

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

func Process(topChanReq chan<- map[string]TopData, errChan chan<- error) {
	client := getClient()
	tResponse := TopResponse{}
	topData := make(map[string]TopData)
	apiConf := conf.ApiConfig[conf.TopApi]
	pageNumInt, err := strconv.Atoi(apiConf.Options[conf.PageParam])
	if err != nil {
		errChan <- err
	}
	var wg sync.WaitGroup
	wg.Add(pageNumInt)
	for i := 0; i < pageNumInt; i++ {
		currentPage := strconv.Itoa(i)
		req, err := tResponse.CreateRequest(apiConf, currentPage)
		if err != nil {
			errChan <- err
		}
		response, err := tResponse.GetResponse(req, client)
		if err != nil {
			errChan <- err
		}
		err = json.Unmarshal(response, &tResponse)
		if err != nil {
			errChan <- err
		}
		err = tResponse.ValidateResponse()
		if err != nil {
			errChan <- err
		}
		for _, v := range tResponse.Data {
			topData[v.CoinInfo.Symbol] = TopData{Symbol: v.CoinInfo.Symbol, Rank: len(topData) + 1}
		}
		wg.Done()
	}
	wg.Wait()
	topChanReq <- topData
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
		return nil, errorsPkg.Wrap(err, "Error during sendig TOP request")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "Error during reading TOP response data")
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
		return nil, errorsPkg.Wrap(err, "Error during creating TOP request")
	}
	req.Header.Add(apiData.CredentialsHeader, apiData.Credentials)
	return req, nil
}

func getClient() http.Client {
	client := http.Client{}
	return client
}
