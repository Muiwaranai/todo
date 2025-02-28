package server

import (
	"fmt"
	"net/http"
	"todo/data"
)

func (s *Server) Greetings(w http.ResponseWriter, r *http.Request) {
	task := data.ToDoTaskDTO{Title: "37", Desctiption: "Nope"}
	s.db.Save(task)
	fmt.Fprintf(w, "Hello from Greetings")
}
