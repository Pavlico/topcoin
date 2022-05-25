package main

import (
	"fmt"
	"log"

	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/top"
	prettifier "github.com/Pavlico/topcoin/services/topcollector/pkg/utils"
)

func main() {
	data, err := top.GetTopData()
	if err != nil {
		log.Println(err)
		return
	}
	prettyData, err := prettifier.PrettyPrint(data)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(prettyData))
}
