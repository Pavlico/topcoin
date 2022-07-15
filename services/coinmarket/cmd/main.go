package main

import (
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/service"
)

func main() {
	if err := service.Run(); err != nil {
		panic(err)
	}
}
