package controller

import (
	"encoding/json"
	"net/http"
	store "rest-api-mid/domain/model"

	"github.com/gorilla/mux"
)

// get controller for GET requests return value by provided key
func GetValueByKey(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	value := store.Get(params["key"])

	if len(value) == 0 {
		json.NewEncoder(response).Encode("Sorry there is no value by this key")
		return
	}

	json.NewEncoder(response).Encode(value)
}

// create controller for POST requests return no data but creates pair in store
func CreatePair(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	key := params["key"]
	value := params["value"]

	if len(key) == 0 {
		json.NewEncoder(response).Encode("Your key is empty")
		return
	}

	store.Set(key, value)
	json.NewEncoder(response).Encode("Created pair: " + key + " and value: " + value)
}

// update controller for PUT requests return no data but updates store by provided key and value
func UpdateValueByKey(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	key := params["key"]
	value := params["value"]

	if len(key) == 0 {
		json.NewEncoder(response).Encode("Your key is empty")
		return
	}

	store.Update(key, value)
	json.NewEncoder(response).Encode("Changed key: " + key + " and value: " + value)
}
