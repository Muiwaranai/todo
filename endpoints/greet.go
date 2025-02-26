package endpoints

import (
	"fmt"
	"net/http"
)

func Greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Greet")
}
