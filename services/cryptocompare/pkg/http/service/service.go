package service

import (
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/dataTypes"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/top"
)

func GetRanks() (map[string]dataTypes.TopData, error) {
	topData, err := top.GetTopData()
	if err != nil {
		return nil, err
	}

	return topData, nil
}
