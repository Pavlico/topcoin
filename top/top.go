package top

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"topcoin/conf"
)

type PrittyResponse struct {
	Name string
	Rank int
}

type TopResponse struct {
	Data []TopResponseData `json:"Data"`
	TopResponseError
}
type TopResponseData struct {
	CoinInfo struct {
		Name string `json:"Name"`
	} `json:"CoinInfo"`
}
type TopResponseError struct {
	Message string `json:"Message"`
}

func Process(client http.Client) ([]byte, error) {
	tResponse := TopResponse{}
	prittyResp := []PrittyResponse{}
	apiConf := conf.ApiConfig[conf.TopApi]
	pageNumInt, err := strconv.Atoi(apiConf.Options[conf.PageParam])
	if err != nil {
		return nil, err
	}
	for i := 0; i < pageNumInt; i++ {
		req, err := tResponse.CreateRequest(apiConf)
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
		if err != nil {
			return nil, err
		}
		for _, v := range tResponse.Data {
			prittyResp = append(prittyResp, PrittyResponse{Name: v.CoinInfo.Name, Rank: len(prittyResp) + 1})
		}
	}
	return PrettyPrint(prittyResp)
}

func (tResponse TopResponse) ValidateResponse() error {
	if tResponse.TopResponseError.Message != conf.SuccessMessage {
		return errors.New("Response is not valid")
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

func (tr TopResponse) CreateRequest(apiData conf.ApiData) (*http.Request, error) {
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

func PrettyPrint(i interface{}) ([]byte, error) {
	return json.MarshalIndent(i, " ", "\t")
}
