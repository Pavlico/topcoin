package service

import (
	"context"
	"log"

	"github.com/Pavlico/topcoin/services/topcollector/pkg/assembler"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/dataTypes"
	errorsPkg "github.com/pkg/errors"
)

func GetCoins() ([]dataTypes.CoinData, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	outputChan := make(chan []dataTypes.CoinData)
	errorChan := make(chan error)
	go assembler.Get(outputChan, errorChan, ctx)
	defer cancel()
	select {
	case err := <-errorChan:
		if errorsPkg.Unwrap(err) != nil {
			err = errorsPkg.Unwrap(err)
		}
		log.Println(err)
		cancel()
		return nil, err
	case v := <-outputChan:

		return v, nil
	}
}
