package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func startPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IgozOverdose")
}

func main() {
	connectionStr := "postgres://api:dota2govno@localhost:5432/apiuser?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println("hui")
		panic(err)
	}
	fmt.Println("Krasava")
	conn.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", startPage)
	log.Panic(http.ListenAndServe(":8080", router))
}
