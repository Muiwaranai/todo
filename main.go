package main

import (
	"net/http"
	"todo/data"
	"todo/server"
)

func main() {
	databse := data.NewDatabase("database/database.json")

	r := http.NewServeMux()
	s := server.NewServerL(*databse)
	r.HandleFunc("/", s.Greetings)

	http.ListenAndServe(":8080", r)
}
