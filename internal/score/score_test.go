package score_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"topcoin/internal/conf"
	"topcoin/internal/score"

	"github.com/gabrielf/gnock"
	"github.com/stretchr/testify/assert"
)

type apiResponseMock struct {
}

func (rm apiResponseMock) GetResponse(req *http.Request, client http.Client) ([]byte, error) {
	transport := gnock.Gnock("https://pro-api.coinmarketcap.com").Get("/v1/cryptocurrency/listings/latest").Reply(200, `{
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
		}`)
	response, _ := transport.RoundTrip(req)
	return ioutil.ReadAll(response.Body)
}
func (rm apiResponseMock) CreateRequest(apiData conf.ApiData) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
}
func TestSuccesfullProcess(t *testing.T) {
	apiResponse := score.Initialize()
	apiResponse = apiResponseMock{}
	responseMock := apiResponse
	scoreStruct := score.ScoreResponse{}
	prettyResp := []score.PrettyResponse{}
	req, err := responseMock.CreateRequest(conf.ApiConfig[conf.ScoreApi])
	resp, err := responseMock.GetResponse(req, http.Client{})
	err = json.Unmarshal(resp, &scoreStruct)
	if err != nil {
		t.Errorf("Should produce error on unmarshal %v", err)
	}
	err = scoreStruct.ValidateResponse(resp)
	if err != nil {
		t.Errorf("Shouldnt produce error on validation %v", err)
	}

	for _, v := range scoreStruct.Data {
		prettyResp = append(prettyResp, score.PrettyResponse{Symbol: v.Symbol, Score: v.Score.Currency.Price})
	}
	assert.Equal(t, prettyResp, []score.PrettyResponse{{Symbol: "BTC", Score: 28379.49}})
}
