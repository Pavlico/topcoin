package handler

import (
	"net/http"

	"github.com/Pavlico/topcoin/internal/service"
	"github.com/Pavlico/topcoin/internal/utils/response"

	"github.com/julienschmidt/httprouter"
)

func GetCoins() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		result, status := service.GetCoins(w, r, ps)
		response.WriteResponse(&w, status, result)
	}
}
