package service

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Pavlico/topcoin/internal/dataTypes"
)

func GetCoins() ([]dataTypes.CoinData, error) {
	coins, err := doRequest()
	if err != nil {
		return nil, err
	}
	return coins, nil

}

func doRequest() ([]dataTypes.CoinData, error) {
	resp, err := http.Get("http://topcollector:8070/topcoins")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	return []dataTypes.CoinData{}, nil
}
