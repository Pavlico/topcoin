package service

import (
	"net/http"

	"github.com/Pavlico/topcoin/internal/coinservice"
	"github.com/Pavlico/topcoin/internal/utils/prettifier"
)

func GetCoins(w http.ResponseWriter, r *http.Request) ([]byte, int) {
	coins, err := coinservice.GetTopCoin()
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	result, err := prettifier.PrettyPrint(coins)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return result, http.StatusOK
}
