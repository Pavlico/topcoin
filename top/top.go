package top

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"topcoin/conf"
	"topcoin/worker"
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

func Process(client http.Client, worker worker.Worker) {
	var tr = TopResponse{}
	var prittyResp = []PrittyResponse{}
	apiData := tr.SetApiData()
	var wg sync.WaitGroup
	wg.Add(conf.AmountOfPages)
	for i := 0; i < conf.AmountOfPages; i++ {
		apiData.Options["page"] = strconv.Itoa(i)
		req := tr.CreateRequest(apiData, worker)
		response := tr.GetResponse(req, client, worker)
		tr.ValidateResponse(response, worker, apiData)
		parsedRep := tr.ParseResponse(response, worker)
		for _, v := range parsedRep.Data {
			prittyResp = append(prittyResp, PrittyResponse{Name: v.CoinInfo.Name, Rank: len(prittyResp) + 1})
		}
		wg.Done()
	}
	wg.Wait()
	fmt.Println(tr.PrettyPrint(prittyResp))
}

func (tr TopResponse) ValidateResponse(response []byte, worker worker.Worker, apiData conf.ApiData) {
	err := json.Unmarshal(response, &tr.TopResponseError)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
		worker.Stop()
	}
	if tr.TopResponseError.Message != conf.SuccessMessage {
		log.Fatalf("Response is not valid Err message: %v", tr.TopResponseError.Message)
		worker.Stop()
	}
}

func (tr TopResponse) ParseResponse(response []byte, worker worker.Worker) TopResponse {
	err := json.Unmarshal(response, &tr)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
		worker.Stop()
	}
	return tr
}

func (tr TopResponse) GetResponse(req *http.Request, client http.Client, worker worker.Worker) []byte {
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
		worker.Stop()
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return responseData
}

func (sr TopResponse) CreateRequest(apiData conf.ApiData, worker worker.Worker) *http.Request {
	param := url.Values{}
	for i, val := range apiData.Options {
		param.Add(i, val)
	}

	endpoint := apiData.ApiAddress + apiData.Subdirectory + param.Encode()
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println(err)
		worker.Stop()
	}
	req.Header.Add("x-api-key", apiData.Credentials)
	return req
}

func (tr TopResponse) SetApiData() conf.ApiData {
	return conf.ApiData{
		ApiAddress:   conf.CryptocompareApiAddress,
		Subdirectory: conf.TopSubDirectory,
		Credentials:  conf.CryptocompareCredentials,
		Options: map[string]string{
			"page":  conf.FirstPage,
			"limit": conf.PageLimit,
			"tsym":  conf.UsdCurrency,
		},
	}
}

func (tr TopResponse) PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, " ", "\t")
	return string(s)
}
