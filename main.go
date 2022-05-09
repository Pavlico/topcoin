package main

import (
	"log"
	"net/http"
	"topcoin/conf"
	"topcoin/flags"
	"topcoin/score"
	"topcoin/top"
	"topcoin/worker"
)

type Api interface {
	ValidateResponse(response []byte, worker worker.Worker, apiData conf.ApiData)
	ParseResponse(response []byte, w worker.Worker)
	CreateRequest(apiData conf.ApiData, w worker.Worker) *http.Request
	GetResponse(req *http.Request, c http.Client, w worker.Worker) []byte
	SetApiData() conf.ApiData
	PrettyPrint(i interface{}) string
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
	worker := worker.NewWorker(flagData.MaxRoutines)
	go processRequest(worker)
	worker.RespChan <- flagData.ApiType
	<-worker.Quit
}

func processRequest(w worker.Worker) {
	for {
		select {
		case reqType := <-w.RespChan:
			if err := get(reqType, w); err != nil {
				log.Println(err)
				w.Stop()
			}

		case <-w.Quit:
			return
		}
	}
}

func get(reqType string, w worker.Worker) error {
	switch reqType {
	case conf.TopApiFlag:
		top.Process(getClient(), w)
		return nil
	case conf.ScoreApiFlag:
		score.Process(getClient(), w)
		return nil
	}
	return nil
}

func getClient() http.Client {
	client := http.Client{}
	return client
}
