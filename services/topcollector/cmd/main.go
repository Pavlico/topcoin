package main

import (
	"fmt"
	"log"

	"github.com/Pavlico/topcoin/services/topcollector/internal/assembler"
	"github.com/Pavlico/topcoin/services/topcollector/internal/flags"
)

func main() {
	flagData := flags.SetFlagData()
	if err := flagData.ValidateFlags(); err != nil {
		log.Println(err)
		return
	}
	data, err := assembler.Get(flagData.ApiTypes)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(data)
}
