package service

import (
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/top"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/utils/prettifier"
)

func GetRanks() ([]byte, error) {
	topData, err := top.GetTopData()
	if err != nil {
		return nil, err
	}
	result, err := prettifier.PrettyPrint(topData)
	if err != nil {
		return nil, err
	}
	return result, nil
}
