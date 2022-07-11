package handler

import (
	"net/http"
	"strings"

	"github.com/Pavlico/topcoin/services/coinmarket/pkg/http/service"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/utils/prettifier"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/utils/response"
)

const symbolParam = "symbols"

const commaSeparator = ","

func GetScores() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		symbolsString := r.FormValue(symbolParam)
		symbols := strings.Split(symbolsString, commaSeparator)
		score, err := service.GetScores(symbols)
		if err != nil {
			status = http.StatusInternalServerError
		}
		result, err := prettifier.PrettyPrint(score)
		if err != nil {
			status = http.StatusInternalServerError
			result = []byte("Internal error")
		}
		response.WriteResponse(&w, status, result)
	}
}
