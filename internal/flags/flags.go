package flags

import (
	"errors"
	"flag"
	"reflect"
)

const emptyFlagValue = ""
const emptyStringVal = ""
const defaultTrheads = 4
const flagLeadApi = "leadapi"
const flagCredentials = "credentials"
const flagHelpMessage = "Some flag message"
const flagRoutines = "routines"

type FlagData struct {
	ApiTypes    []string
	LeadApi     string
	MaxRoutines int
}

func SetFlagData() FlagData {
	routines := flag.Int(flagRoutines, defaultTrheads, flagHelpMessage)
	leadApi := flag.String(flagLeadApi, emptyFlagValue, flagHelpMessage)
	flag.Parse()
	return FlagData{
		ApiTypes:    flag.Args(),
		MaxRoutines: *routines,
		LeadApi:     *leadApi,
	}
}

func (fd FlagData) ValidateFlags() error {
	flagValues := reflect.ValueOf(fd)
	for i := 0; i < flagValues.NumField(); i++ {
		if flagValues.Field(i).Interface() == emptyFlagValue {
			return errors.New("You are missing flag values")
		}
	}
	return nil
}
