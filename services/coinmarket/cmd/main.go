package main

import (
	"fmt"
	"log"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/flags"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/score"
	prettifier "github.com/Pavlico/topcoin/services/topcollector/pkg/utils"
)

func main() {
	symbols := flags.GetFlagSymbols()
	data, err := score.GetScoreData(symbols)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(prettifier.PrettyPrint(data))
}
