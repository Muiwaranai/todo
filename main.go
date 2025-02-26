package main

import (
	"net/http"
	"todo/modules"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", modules.Greetings)
	r.HandleFunc("/", modules.Greetings)
	r.HandleFunc("/", modules.Greetings)
	r.HandleFunc("/", modules.Greetings)
	r.HandleFunc("/", modules.Greetings)

	http.ListenAndServe(":8080", r)
}
