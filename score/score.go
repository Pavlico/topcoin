package score

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"topcoin/conf"
)

type PrittyResponse struct {
	Name  string
	Score float32
}

type ScoreResponse struct {
	Data []ScoreResponseData `json:"Data"`
	ScoreResponseError
}
type ScoreResponseData struct {
	Name  string `json:"Name"`
	Score struct {
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

func Process(client http.Client) ([]byte, error) {
	var sr = ScoreResponse{}
	var prittyResp = []PrittyResponse{}
	var apiConf = conf.ApiConfig[conf.ScoreApi]
	req, err := sr.CreateRequest(apiConf)
	if err != nil {
		return nil, err
	}
	response, err := sr.GetResponse(req, client)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &sr)
	if err != nil {
		return nil, err
	}
	err = sr.ValidateResponse(response)
	if err != nil {
		return nil, err
	}
	parsedRep, err := sr.ParseResponse(response)
	if err != nil {
		return nil, err
	}
	for _, v := range parsedRep.Data {
		prittyResp = append(prittyResp, PrittyResponse{Name: v.Name, Score: v.Score.Currency.Price})
	}
	return PrettyPrint(prittyResp)
}

func (sr ScoreResponse) GetResponse(req *http.Request, client http.Client) ([]byte, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return responseData, nil
}
func (sr ScoreResponse) ValidateResponse(response []byte) error {

	if sr.ScoreResponseError.Status.ErrorCode != conf.NoErrorCode {
		return errors.New("Response is not valid")
	}
	return nil
}

func (sr *ScoreResponse) ParseResponse(response []byte) (*ScoreResponse, error) {
	err := json.Unmarshal(response, &sr)
	if err != nil {
		return nil, err
	}
	return sr, nil
}

func (sr ScoreResponse) CreateRequest(apiData conf.ApiData) (*http.Request, error) {

	endpoint := apiData.ApiAddress + apiData.EndPoint
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(apiData.CredentialsHeader, apiData.Credentials)
	return req, nil
}

func PrettyPrint(i interface{}) ([]byte, error) {
	return json.MarshalIndent(i, "", "\t")
}
