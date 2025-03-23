package users

import (
	"fmt"
	"net/http"
)



func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("endpoint creation nouvel utilisateur")
}

