package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

const emptyFlagValue = ""
const emptyStringVal = ""
const defaultTrheads = 4
const flagApi = "api"
const flagCredentials = "credentials"
const flagHelpMessage = "Some flag message"
const flagRoutines = "routines"
const usdCurrency = "USD"
const cryptocompareCredentials = "93d94cc6b38c0c4ed61c744e74d9ecd14223f427b14aea9e7f68f9ca72a5ab78"
const coinmarketcapCredentials = "18f09716-386e-440a-8906-551d607d1574"
const cryptocompareApiAddress = "https://min-api.cryptocompare.com"
const coinmarketcapAddress = "https://pro-api.coinmarketcap.com"
const topSubDirectory = "/data/top/totalvolfull?"
const scoreSubDirectory = "/v1/cryptocurrency/listings/latest?"
const topApiFlag = "top"
const scoreApiFlag = "score"
const amountOfPages = 2
const pageLimit = "5"
const firstPage = "0"
const successMessage = "Success"

type apiData struct {
	apiAddress   string
	subdirectory string
	credentials  string
	options      map[string]string
	response     interface{}
}

type FlagData struct {
	ApiType string
	// ApiCredentials string
	MaxRoutines int
}
type TopResponse struct {
	Data []TopResponseData `json:"Data"`
	TopResponseError
}
type TopResponseData struct {
	CoinInfo struct {
		Name string `json:"Name"`
		Rank int    `json:"Rank"`
	} `json:"CoinInfo"`
}
type TopResponseError struct {
	Response string `json:"Response"`
}

type ParsedResp struct {
	Rank string
	Name string
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

type Worker struct {
	RespChan chan apiData
	Quit     chan bool
}

func main() {
	apis := getApiMap()
	flagData := setFlagData()
	if err := flagData.validateFlags(); err != nil {
		log.Println(err)
		return
	}
	worker := NewWorker(flagData.MaxRoutines)
	go worker.processRequest()
	_, ok := apis[flagData.ApiType]
	if ok {
		worker.RespChan <- apis[flagData.ApiType]
	}
	<-worker.Quit
}

func NewWorker(maxRoutines int) Worker {
	return Worker{
		RespChan: make(chan apiData, maxRoutines),
		Quit:     make(chan bool),
	}
}

func (w Worker) processRequest() {
	for {
		select {
		case apiData := <-w.RespChan:
			if err := w.get(apiData); err != nil {
				log.Println(err)
				w.Stop()
			}

		case <-w.Quit:
			return
		}
	}
}

func (w Worker) newRequest(apiData apiData) *http.Request {
	param := url.Values{}
	for i, val := range apiData.options {
		param.Add(i, val)
	}

	endpoint := apiData.apiAddress + apiData.subdirectory + param.Encode()
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println(err)
		w.Stop()
	}
	req.Header.Add("x-api-key", apiData.credentials)
	req.Header.Add("X-CMC_PRO_API_KEY", apiData.credentials)
	return req
}

func (w Worker) get(apiData apiData) error {
	req := w.newRequest(apiData)
	response := getResponse(req, w)
	validateResponse(response, w, apiData)
	parsedRep := parseResponse(response, w, apiData)
	fmt.Println(PrettyPrint(parsedRep))
	return nil
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
func getResponse(req *http.Request, w Worker) []byte {
	client := getClient()
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to API endpoint. %+v", err)
		w.Quit <- true
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return responseData
}

func validateResponse(response []byte, w Worker, apiData apiData) {
	// err := json.Unmarshal(response, &responseType)
	// if err != nil {
	// 	log.Fatalf("Couldn't parse response body. %+v", err)
	// 	w.Quit <- true
	// }
	// _, okTop := responseType.(TopResponseError)
	// _, okScore := responseType.(ScoreResponseError)
	// if responseType.Message != successMessage {
	// 	log.Fatalf("Response is not valid %+v", err)
	// 	w.Quit <- true
	// }
}

func parseResponse(response []byte, w Worker, apiData apiData) interface{} {
	respStruct := apiData.response
	err := json.Unmarshal(response, &respStruct)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
		w.Quit <- true
	}
	return respStruct
}

func setFlagData() FlagData {
	apiType := flag.String(flagApi, emptyFlagValue, flagHelpMessage)
	routines := flag.Int(flagRoutines, defaultTrheads, flagHelpMessage)
	flag.Parse()
	return FlagData{
		ApiType:     *apiType,
		MaxRoutines: *routines,
	}
}

func (fd FlagData) validateFlags() error {
	flagValues := reflect.ValueOf(fd)
	for i := 0; i < flagValues.NumField(); i++ {
		if flagValues.Field(i).Interface() == emptyFlagValue {
			return errors.New("You are missing flag values")
		}
	}
	return nil
}

func getClient() http.Client {
	client := http.Client{}
	return client
}

func getApiMap() map[string]apiData {
	apiMap := map[string]apiData{
		topApiFlag: {
			apiAddress:   cryptocompareApiAddress,
			subdirectory: topSubDirectory,
			credentials:  cryptocompareCredentials,
			options: map[string]string{
				"page":  firstPage,
				"limit": pageLimit,
				"tsym":  "USD",
			},
			response: TopResponse{},
		},
		scoreApiFlag: {
			apiAddress:   coinmarketcapAddress,
			subdirectory: scoreSubDirectory,
			credentials:  coinmarketcapCredentials,
			response:     ScoreResponse{},
		},
	}
	return apiMap
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}
