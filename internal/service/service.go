package service

import (
	"net/http"

	"github.com/Pavlico/topcoin/internal/coinservice"
	"github.com/julienschmidt/httprouter"
)

func GetCoins(w http.ResponseWriter, r *http.Request, ps httprouter.Params) ([]byte, int) {
	result, err := coinservice.GetTopCoin()
	if err != nil {
		return result, http.StatusInternalServerError
	}
	return result, http.StatusOK
}
