package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func Greetings(w http.ResponseWriter, r *http.Request) {
	var users []Person

	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if len(users) == 0 {
		http.Error(w, "Empty body", http.StatusBadRequest)
	}

	fmt.Printf("Name: %s, Age: %d, Length: %d\n", users[0].Name, users[0].Age, len(users))

	fmt.Fprintf(w, "Received user: %+v", users)
}
