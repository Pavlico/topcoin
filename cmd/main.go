package main

import (
	"fmt"
	"log"
	"topcoin/internal/assembler"
	"topcoin/internal/flags"
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
