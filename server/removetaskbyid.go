package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) RemoveTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	s.db.RemoveById(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task rmoved from database")
}
