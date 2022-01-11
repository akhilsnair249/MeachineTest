package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RespondFailure(w http.ResponseWriter, statusCode int, errorObj error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorObj)

	log.Println("Request Failed :", statusCode, " ,", errorObj)
}

func RespondSuccess(w http.ResponseWriter, responseObj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseObj)
}

func GetID(r *http.Request) string {
	vars := mux.Vars(r)
	idString := vars["id"]
	return idString
}
