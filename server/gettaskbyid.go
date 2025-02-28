package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) GetTaskById(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	if queryId == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(queryId)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := s.db.GetById(id)
	if err != nil {
		http.Error(w, "Not task not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Failed to encode task", http.StatusInternalServerError)
	}
}
