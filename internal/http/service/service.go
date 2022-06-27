package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Pavlico/topcoin/internal/dataTypes"
	"github.com/Pavlico/topcoin/internal/utils/prettifier"
	errorsPkg "github.com/pkg/errors"
)

func GetCoins() ([]byte, int) {
	coins, err := doRequest()
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	result, err := prettifier.PrettyPrint(coins)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return result, http.StatusOK
}

func doRequest() ([]dataTypes.CoinData, error) {
	client := getClient()
	req, err := createRequest("http://topcollector:8070/topcoins")
	if err != nil {
		return nil, err
	}
	responseData := []dataTypes.CoinData{}
	response, err := getResponse(req, client)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
func createRequest(endpoint string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	return req, nil

}

func getResponse(req *http.Request, client http.Client) ([]byte, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, errorsPkg.Wrap(err, "Error during sending request")
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
