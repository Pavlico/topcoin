package flags

import (
	"errors"
	"flag"
	"reflect"
)

const emptyFlagValue = ""
const emptyStringVal = ""
const defaultTrheads = 4
const defaultType = "default"
const flagCredentials = "credentials"
const flagHelpMessage = "Some flag message"
const flagRoutines = "routines"
const flagType = "type"

type FlagData struct {
	ApiTypes    []string
	MaxRoutines int
	RequestType string
}

func GetFlagData() FlagData {
	routines := flag.Int(flagRoutines, defaultTrheads, flagHelpMessage)
	requestType := flag.String(flagType, defaultType, flagHelpMessage)
	flag.Parse()
	return FlagData{
		ApiTypes:    flag.Args(),
		MaxRoutines: *routines,
		RequestType: *requestType,
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
