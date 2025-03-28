package lines

import (
	"fmt"
	"net/http"
)


func GetTransportModesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("endpoint récupérant les modes de transports")
}

func GetTransportModeLinesIdsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("endpoint récupérant les ids des lignes du mode de transport")
}