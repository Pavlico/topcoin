package handler

import (
	"net/http"

	"github.com/Pavlico/topcoin/internal/utils/prettifier"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/http/service"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/utils/response"
)

func GetMergedData() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		coins, err := service.GetCoins()
		status := http.StatusOK
		if err != nil {
			status = http.StatusInternalServerError
		}
		result, err := prettifier.PrettyPrint(coins)
		if err != nil {
			status = http.StatusInternalServerError
			result = []byte("Internal error")
		}
		response.WriteResponse(&w, status, result)
	}
}
