package main

import (
	"fmt"
	"log"
	"net/http"
	"topcoin/conf"
	"topcoin/flags"
	"topcoin/score"
	"topcoin/top"
)

type Api interface {
	CreateRequest(apiData conf.ApiData) (*http.Request, error)
	GetResponse(req *http.Request, client http.Client) ([]byte, error)
}

type ParsedResp struct {
	Rank string
	Name string
}

func main() {
	flagData := flags.SetFlagData()
	if err := flagData.ValidateFlags(); err != nil {
		log.Println(err)
		return
	}
	response, err := get(flagData.ApiType)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(response))
}

func get(reqType string) ([]byte, error) {
	switch reqType {
	case conf.TopApi:
		return top.Process(getClient())
	case conf.ScoreApi:
		return score.Process(getClient())
	}
	return nil, nil
}

func getClient() http.Client {
	client := http.Client{}
	return client
}
