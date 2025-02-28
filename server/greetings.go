package server

import (
	"fmt"
	"net/http"
)

func (s *Server) Greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Greetings")
}
