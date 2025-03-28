package auth

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("endpoint permettant à l'utilisateur de se connecter")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("endpoint permettant à l'utilisateur de se connecter")
}
