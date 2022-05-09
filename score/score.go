package score

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"topcoin/conf"
	"topcoin/worker"
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

func Process(client http.Client, worker worker.Worker) {
	var sr = ScoreResponse{}
	var prittyResp = []PrittyResponse{}
	apiData := sr.SetApiData()
	req := sr.CreateRequest(apiData, worker)
	response := sr.GetResponse(req, client, worker)
	sr.ValidateResponse(response, worker)
	parsedRep := sr.ParseResponse(response, worker)
	for _, v := range parsedRep.Data {
		prittyResp = append(prittyResp, PrittyResponse{Name: v.Name, Score: v.Score.Currency.Price})
	}
	fmt.Println(sr.PrettyPrint(prittyResp))
}

func (sr ScoreResponse) GetResponse(req *http.Request, client http.Client, worker worker.Worker) []byte {
	response, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to API endpoint. %+v", err)
		worker.Quit <- true
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return responseData
}
func (sr ScoreResponse) ValidateResponse(response []byte, worker worker.Worker) {
	err := json.Unmarshal(response, &sr.ScoreResponseError)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
		worker.Quit <- true
	}
	if sr.ScoreResponseError.Status.ErrorCode != conf.NoErrorCode {
		log.Fatalf("Response is not valid Err code: %v", sr.ScoreResponseError.Status.ErrorCode)
		worker.Quit <- true
	}
}

func (sr ScoreResponse) ParseResponse(response []byte, w worker.Worker) ScoreResponse {
	err := json.Unmarshal(response, &sr)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
		w.Quit <- true
	}
	return sr
}

func (sr ScoreResponse) CreateRequest(apiData conf.ApiData, w worker.Worker) *http.Request {

	endpoint := apiData.ApiAddress + apiData.Subdirectory
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println(err)
		w.Stop()
	}
	req.Header.Add("X-CMC_PRO_API_KEY", apiData.Credentials)
	return req
}

func (sr ScoreResponse) SetApiData() conf.ApiData {
	return conf.ApiData{
		ApiAddress:   conf.CoinmarketcapAddress,
		Subdirectory: conf.ScoreSubDirectory,
		Credentials:  conf.CoinmarketcapCredentials,
	}
}

func (sr ScoreResponse) PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
