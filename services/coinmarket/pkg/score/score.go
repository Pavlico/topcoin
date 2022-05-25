package score

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/conf"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/dataTypes"

	errorsPkg "github.com/pkg/errors"
)

func GetScoreData(symbols []string) (map[string]dataTypes.ScoreData, error) {
	client := getClient()
	sResponse := dataTypes.ScoreResponse{}
	scoreData := make(map[string]dataTypes.ScoreData)
	apiConf := conf.ApiConfig[conf.ScoreApi]
	req, err := CreateRequest(apiConf, symbols)
	if err != nil {
		return nil, err
	}
	response, err := GetResponse(req, client)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &sResponse)
	if err != nil {
		return nil, err
	}
	err = ValidateResponse(sResponse)
	if err != nil {
		return nil, err
	}
	for _, v := range sResponse.Data {
		scoreData[v.Symbol] = dataTypes.ScoreData{Symbol: v.Symbol, Score: v.Score.Currency.Price}
	}
	return scoreData, nil
}

func GetResponse(req *http.Request, client http.Client) ([]byte, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "Error during sendig SCORE request")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "Error during reading SCORE response data")
	}
	defer response.Body.Close()
	return responseData, nil
}
func ValidateResponse(sResponse dataTypes.ScoreResponse) error {
	if sResponse.ScoreResponseError.Status.ErrorCode != conf.NoErrorCode {
		return errors.New("Response SCORE not valid")
	}
	return nil
}

func CreateRequest(apiData conf.ApiData, symbols []string) (*http.Request, error) {
	param := url.Values{}
	for i, val := range apiData.Options {
		param.Add(i, val)
	}
	param.Add(conf.SymbolParam, strings.Join(symbols, ","))
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
		Timeout: conf.ApiTimeout * time.Second,
	}
	return client
}
