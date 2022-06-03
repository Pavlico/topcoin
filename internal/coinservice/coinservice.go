package coinservice

import (
	"context"

	"github.com/Pavlico/topcoin/services/topcollector/pkg/assembler"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/dataTypes"
	errorsPkg "github.com/pkg/errors"
)

func GetTopCoin() ([]dataTypes.CoinData, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	outputChan := make(chan []dataTypes.CoinData)
	errorChan := make(chan error)
	defer cancel()
	go assembler.Get(outputChan, errorChan, ctx)
	select {
	case err := <-errorChan:
		if errorsPkg.Unwrap(err) != nil {
			err = errorsPkg.Unwrap(err)
		}
		return nil, err
	case v := <-outputChan:
		return v, nil
	}
}
