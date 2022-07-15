package main

import (
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/service"
)

func main() {
	if err := service.Run(); err != nil {
		panic(err)
	}
}
