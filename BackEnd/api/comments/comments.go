package comments

import (
	"fmt"
	"net/http"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("endpoint permettant à l'utilisateur de rajouter des commentaires")
}

func GetEventCommentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("endpoint donnant les commentaires d'un évenement")
}
