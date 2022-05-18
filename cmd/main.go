package main

import (
	"context"
	"log"
	"topcoin/internal/assembler"
	"topcoin/internal/dataTypes"
	"topcoin/internal/flags"
	prettifier "topcoin/internal/utils"

	errorsPkg "github.com/pkg/errors"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	outputChan := make(chan []dataTypes.CoinData)
	errorChan := make(chan error)
	flagData := flags.SetFlagData()
	if err := flagData.ValidateFlags(); err != nil {
		log.Println(err)
		return
	}
	go assembler.Get(flagData.ApiTypes, outputChan, errorChan, ctx)
	select {
	case err := <-errorChan:
		if errorsPkg.Unwrap(err) != nil {
			err = errorsPkg.Unwrap(err)
		}
		log.Println(err)
		cancel()
		return
	case v := <-outputChan:
		prettyResp, err := prettifier.PrettyPrint(v)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(prettyResp))
		return
	}
}
