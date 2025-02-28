package server

import (
	"net/http"
)

func (s *Server) Greetings(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/instructions.html")
}
