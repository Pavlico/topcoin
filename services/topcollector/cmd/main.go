package main

import (
	grpcserver "github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/server"
	httpserver "github.com/Pavlico/topcoin/services/topcollector/pkg/http/server"
)

const grpcType = "grpc"
const defaultType = "default"

func main() {
	if true {
		serverManager := httpserver.InitServer()
		serverManager.Serve()
	}
	if false {
		grpcserver.Serve()
	}
	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// outputChan := make(chan []dataTypes.CoinData)
	// errorChan := make(chan error)
	// flagData := flags.GetFlagData()
	// if flagData.RequestType == grpcType {
	// 	go grpcshandler.NewCoinList().GetTopCoins(ctx, &protos.TopRequest{})
	// }
	// if flagData.RequestType == defaultType {
	// 	go assembler.Get(outputChan, errorChan, ctx)
	// }
	// select {
	// case err := <-errorChan:
	// 	if errorsPkg.Unwrap(err) != nil {
	// 		err = errorsPkg.Unwrap(err)
	// 	}
	// 	log.Println(err)
	// 	cancel()
	// 	return
	// case v := <-outputChan:
	// 	prettyResp, err := prettifier.PrettyPrint(v)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	log.Println(string(prettyResp))
	// 	return
	// }
}
