package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connectionStr := "postgres://api:dota2govno@localhost:5432/apiuser?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println("hui")
		panic(err)
	}
	fmt.Println("Krasava")
	conn.Close()

}
