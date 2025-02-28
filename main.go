package main

import (
	"net/http"
	"todo/data"
	"todo/server"
)

func main() {
	databse := data.NewDatabase()

	r := http.NewServeMux()
	s := server.NewServerL(*databse)
	r.HandleFunc("/", s.Greetings)
	r.HandleFunc("/api/getall/", s.GetAllTasks)
	r.HandleFunc("/api/addtask", s.AddTask)
	r.HandleFunc("/api/get", s.GetTaskById)
	r.HandleFunc("/api/remove", s.RemoveTaskById)

	http.ListenAndServe(":8080", r)
}
