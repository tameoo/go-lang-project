package router

import (
	store "rest-api-mid/domain/model"
	controller "rest-api-mid/interface/controller"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	// router entity by mux
	router := mux.NewRouter()

	// mock pair create
	store.Create()

	// router methods and controller linking
	router.HandleFunc("/store/{key}", controller.GetValueByKey).Methods("GET")
	router.HandleFunc("/store/{key}/{value}", controller.CreatePair).Methods("POST")
	router.HandleFunc("/store/{key}/{value}", controller.UpdateValueByKey).Methods("PUT")

	return router
}
