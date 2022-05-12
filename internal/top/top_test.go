package top_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"topcoin/internal/conf"
	"topcoin/internal/top"

	"github.com/gabrielf/gnock"
	"github.com/stretchr/testify/assert"
)

func GetSuccessResponse(req *http.Request, client http.Client) ([]byte, error) {
	transport := gnock.Gnock("https://min-api.cryptocompare.com").Get("/data/top/totalvolfull").Reply(200,
		`{
		"Message": "Success",
		"Data": [
			{
				"CoinInfo": {
					"Name": "BTC"
				}
			},
			{
				"CoinInfo": {
					"Name": "TEST"
				}
			},
			{
				"CoinInfo": {
					"Name": "MEGACOIN"
				}
			}
			]
		}`)
	response, _ := transport.RoundTrip(req)
	return ioutil.ReadAll(response.Body)
}
func CreateRequest(apiData conf.ApiData) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "https://min-api.cryptocompare.com/data/top/totalvolfull?limit=3&tsym=USD", nil)
}
func TestSuccesfullProcess(t *testing.T) {
	topStruct := top.TopResponse{}
	prettyResp := []top.PrettyResponse{}
	req, err := CreateRequest(conf.ApiConfig[conf.ScoreApi])
	resp, err := GetSuccessResponse(req, http.Client{})
	err = json.Unmarshal(resp, &topStruct)
	if err != nil {
		t.Errorf("Shouldnt produce error on unmarshal %v", err)
	}
	err = topStruct.ValidateResponse()
	if err != nil {
		t.Errorf("Shouldnt produce error on validation %v", err)
	}

	for _, v := range topStruct.Data {
		log.Println(v)
		prettyResp = append(prettyResp, top.PrettyResponse{Symbol: v.CoinInfo.Symbol, Rank: len(prettyResp) + 1})
	}
	assert.Equal(t, prettyResp, []top.PrettyResponse{
		{Symbol: "BTC", Rank: 1},
		{Symbol: "TEST", Rank: 2},
		{Symbol: "MEGACOIN", Rank: 3}})
}
