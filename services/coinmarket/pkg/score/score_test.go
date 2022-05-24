package score_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const properResponse = `{
	"status": {
		"error_code": 0,
		"error_message": null
	},
	"data": [
		{
			"name": "Bitcoin",
			"symbol": "BTC",
			"quote": {
				"USD": {
					"price": 28379.490491632012
				}
			}
		}
		]
	}`

func createSuccesfullResponseMock() {
	gock.New("https://pro-api.coinmarketcap.com").Get("/v1/cryptocurrency/listings/latest").Reply(200).JSON(properResponse)
}
func createRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
}
func TestSuccesfullProcess(t *testing.T) {
	scoreStruct := score.ScoreResponse{}
	prettyResp := []dataTypes.ScoreData{}
	req, err := createRequest()
	createSuccesfullResponseMock()
	resp, err := scoreStruct.GetResponse(req, http.Client{})
	err = json.Unmarshal(resp, &scoreStruct)
	if err != nil {
		t.Errorf("Shouldnt produce error on unmarshal %v", err)
	}
	err = scoreStruct.ValidateResponse(resp)
	if err != nil {
		t.Errorf("Shouldnt produce error on validation %v", err)
	}

	for _, v := range scoreStruct.Data {
		prettyResp = append(prettyResp, dataTypes.ScoreData{Symbol: v.Symbol, Score: v.Score.Currency.Price})
	}
	assert.Equal(t, prettyResp, []dataTypes.ScoreData{{Symbol: "BTC", Score: 28379.49}})
}
