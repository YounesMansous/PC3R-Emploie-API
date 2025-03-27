package events

import (
	"fmt"
	"net/http"
)

func GetLineEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("endpoint donnant les évenements d'une ligne de transport")
}

func GetEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("endpoint retournant un évènement")
}
