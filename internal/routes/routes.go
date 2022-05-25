package routes

import (
	"github.com/Pavlico/topcoin/internal/handler"
	"github.com/julienschmidt/httprouter"
)

func GetAvailableRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/top", handler.GetCoins())
	return router
}
