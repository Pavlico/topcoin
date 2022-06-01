package handler

import (
	"net/http"

	"github.com/Pavlico/topcoin/internal/service"
	"github.com/Pavlico/topcoin/internal/utils/response"
)

func GetCoins() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, status := service.GetCoins(w, r)
		response.WriteResponse(&w, status, result)
	}
}
