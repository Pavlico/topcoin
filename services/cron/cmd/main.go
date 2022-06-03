package main

import (
	"context"
	"time"

	"github.com/Pavlico/topcoin/services/dbservice/pkg/database"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/assembler"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/dataTypes"
)

const INTERVAL_PERIOD time.Duration = 2 * time.Hour

func main() {
	jobChan := make(chan bool)
	go runCron(jobChan)
	for {
		<-jobChan
		coinData, err := getTopData()
		if err != nil {
			//place for log
		}
		db, err := database.Initialize()
		if err != nil {
			//place for log
		}
		err = db.Save(coinData)
		if err != nil {
			//place for log
		}
		//log saved data to db
	}
}

func runCron(jobChan chan<- bool) {
	for range time.Tick(INTERVAL_PERIOD) {
		jobChan <- true
	}

}

func getTopData() ([]dataTypes.CoinData, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	outputChan := make(chan []dataTypes.CoinData)
	errorChan := make(chan error)
	defer cancel()
	go assembler.Get(outputChan, errorChan, ctx)
	select {
	case err := <-errorChan:
		return nil, err
	case v := <-outputChan:
		return v, nil
	}
}
