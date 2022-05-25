package main

import (
	"fmt"
	"log"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/conf"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/flags"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
	prettifier "github.com/Pavlico/topcoin/services/topcollector/pkg/utils"
)

func main() {
	symbols := flags.GetFlagSymbols()
	if symbols[0] == conf.EmptyValue {
		fmt.Println("No symbols")
		return
	}
	data, err := score.GetScoreData(symbols)
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
