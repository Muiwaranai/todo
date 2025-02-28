package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := s.db.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
