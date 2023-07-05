package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/registration", RegisterUser).Methods("POST")
	router.HandleFunc("/complete-registration", CompleteRegistration).Methods("PATCH")
	return router
}
