package handler

import (
	"net/http"
	"strings"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/http/service"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/utils/response"
)

const symbolParam = "symbols"

const commaSeparator = ","

func GetScores() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		symbolsString := r.FormValue(symbolParam)
		symbols := strings.Split(symbolsString, commaSeparator)
		result, err := service.GetScores(symbols)
		status := http.StatusOK
		if err != nil {
			status = http.StatusInternalServerError
		}
		response.WriteResponse(&w, status, result)
	}
}
