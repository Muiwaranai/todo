package main

import (
	"net/http"
	"todo/endpoints"
	"todo/jsondatabase"
)

func main() {
	databse := jsondatabase.NewJsonDB("jsondatabase/db.json")

	r := http.NewServeMux()
	r.HandleFunc("/", endpoints.Greetings)
	r.HandleFunc("/api/add", endpoints.Addtask)
	// r.HandleFunc("/", modules.Greetings)
	// r.HandleFunc("/", modules.Greetings)
	// r.HandleFunc("/", modules.Greetings)

	http.ListenAndServe(":8080", r)
}
