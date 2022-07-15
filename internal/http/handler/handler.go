package handler

import (
	"net/http"

	grpcService "github.com/Pavlico/topcoin/internal/grpc/service"
	httpService "github.com/Pavlico/topcoin/internal/http/service"
	"github.com/Pavlico/topcoin/internal/utils/prettifier"
	"github.com/Pavlico/topcoin/internal/utils/response"
)

func GetCoinsHttp() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		coins, status := httpService.GetCoins()
		result, err := prettifier.PrettyPrint(coins)
		if err != nil {
			status = http.StatusInternalServerError
			result = []byte("Internal error")
		}
		response.WriteResponse(&w, status, result)
	}
}

func GetCoinsGrpc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		coins, status := grpcService.GetCoins()
		result, err := prettifier.PrettyPrint(coins)
		if err != nil {
			status = http.StatusInternalServerError
			result = []byte("Internal error")
		}
		response.WriteResponse(&w, status, result)
	}
}
