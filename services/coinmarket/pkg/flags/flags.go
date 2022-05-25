package flags

import (
	"flag"
	"strings"
)

const emptyFlagValue = ""
const flagHelpMessage = "Some flag message"
const flagSymbols = "symbols"
const symbolSplitChar = ","

func GetFlagSymbols() []string {
	symbolsString := flag.String(flagSymbols, emptyFlagValue, flagHelpMessage)
	flag.Parse()
	symbols := strings.Split(*symbolsString, symbolSplitChar)

	return symbols
}
