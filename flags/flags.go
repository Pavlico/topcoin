package flags

import (
	"errors"
	"flag"
	"reflect"
)

const emptyFlagValue = ""
const emptyStringVal = ""
const defaultTrheads = 4
const flagApi = "api"
const flagCredentials = "credentials"
const flagHelpMessage = "Some flag message"
const flagRoutines = "routines"

type FlagData struct {
	ApiType     string
	MaxRoutines int
}

func SetFlagData() FlagData {
	apiType := flag.String(flagApi, emptyFlagValue, flagHelpMessage)
	routines := flag.Int(flagRoutines, defaultTrheads, flagHelpMessage)
	flag.Parse()
	return FlagData{
		ApiType:     *apiType,
		MaxRoutines: *routines,
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
