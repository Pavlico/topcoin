package handler

import (
	"net/http"

	"github.com/Pavlico/topcoin/internal/http/service"
	"github.com/Pavlico/topcoin/internal/utils/response"
)

func GetCoins() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, status := service.GetCoins()
		response.WriteResponse(&w, status, result)
	}
}
