package main

import (
	"net/http"

	"github.com/AlexSparr0w/mmrapi/controllers"
	"github.com/AlexSparr0w/mmrapi/model"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	model.ConnectDatabase()

	server.ListenAndServe()
}
