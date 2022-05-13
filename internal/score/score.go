package score

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"topcoin/internal/conf"
)

type Api interface {
	CreateRequest(apiData conf.ApiData) (*http.Request, error)
	GetResponse(req *http.Request, client http.Client) ([]byte, error)
}

type PrettyResponse struct {
	Symbol string
	Score  float32
}

type ScoreResponse struct {
	Data []ScoreResponseData `json:"Data"`
	ScoreResponseError
}
type ScoreResponseData struct {
	Symbol string `json:"symbol"`
	Score  struct {
		Currency struct {
			Price float32 `json:"price"`
		} `json:"USD"`
	} `json:"quote"`
}

type ScoreResponseError struct {
	Status struct {
		ErrorCode int `json:"error_code"`
	} `json:"status"`
}

func Initialize() Api {
	var apiResponse Api = ScoreResponse{}
	return apiResponse
}

func Process() ([]PrettyResponse, error) {
	client := getClient()
	sResponse := ScoreResponse{}
	prettyResp := []PrettyResponse{}
	apiConf := conf.ApiConfig[conf.ScoreApi]
	req, err := sResponse.CreateRequest(apiConf)
	if err != nil {
		return nil, err
	}
	response, err := sResponse.GetResponse(req, client)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &sResponse)
	if err != nil {
		return nil, err
	}
	err = sResponse.ValidateResponse(response)
	if err != nil {
		return nil, err
	}
	for _, v := range sResponse.Data {
		prettyResp = append(prettyResp, PrettyResponse{Symbol: v.Symbol, Score: v.Score.Currency.Price})
	}
	return prettyResp, nil
}

func (sResponse ScoreResponse) GetResponse(req *http.Request, client http.Client) ([]byte, error) {
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
func (sResponse ScoreResponse) ValidateResponse(response []byte) error {

	if sResponse.ScoreResponseError.Status.ErrorCode != conf.NoErrorCode {
		return errors.New("Response is not valid")
	}
	return nil
}

func (sResponse ScoreResponse) CreateRequest(apiData conf.ApiData) (*http.Request, error) {
	param := url.Values{}
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
	client := http.Client{}
	return client
}
