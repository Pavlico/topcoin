package top_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Pavlico/topcoin/services/cryptocompare/internal/dataTypes"
	"github.com/Pavlico/topcoin/services/cryptocompare/internal/top"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

const properResponse = `{
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
	}`

func createSuccesfullResponseMock() {
	gock.New("https://min-api.cryptocompare.com").Get("/data/top/totalvolfull").Reply(200).JSON(properResponse)

}
func createRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "https://min-api.cryptocompare.com/data/top/totalvolfull?limit=3&tsym=USD", nil)
}
func TestSuccesfullProcess(t *testing.T) {
	defer gock.Off()
	topStruct := top.TopResponse{}
	topData := make(map[string]dataTypes.TopData)
	createSuccesfullResponseMock()
	req, err := createRequest()
	resp, err := topStruct.GetResponse(req, http.Client{})
	err = json.Unmarshal(resp, &topStruct)
	if err != nil {
		t.Errorf("Shouldnt produce error on unmarshal %v", err)
	}
	err = topStruct.ValidateResponse()
	if err != nil {
		t.Errorf("Shouldnt produce error on validation %v", err)
	}

	for _, v := range topStruct.Data {
		topData[v.CoinInfo.Symbol] = dataTypes.TopData{Symbol: v.CoinInfo.Symbol, Rank: len(topData) + 1}
	}

	st.Expect(t, topData, map[string]dataTypes.TopData{
		"BTC":      {Symbol: "BTC", Rank: 1},
		"TEST":     {Symbol: "TEST", Rank: 2},
		"MEGACOIN": {Symbol: "MEGACOIN", Rank: 3}})

	st.Expect(t, gock.IsDone(), true)
}
