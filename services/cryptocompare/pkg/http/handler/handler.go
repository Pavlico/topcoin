package handler

import (
	"net/http"

	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/http/service"
	"github.com/Pavlico/topcoin/services/cryptocompare/pkg/utils/response"
)

func GetRanks() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := service.GetRanks()
		status := http.StatusOK
		if err != nil {
			status = http.StatusInternalServerError
		}
		response.WriteResponse(&w, status, result)
	}
}
