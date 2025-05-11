package comments

import (
	"api/controllers/auth"
	"api/models"
	"api/utils/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")

	var comment models.Comments

	event_id, err := strconv.ParseInt(r.URL.Query().Get("event_id"), 10, 64)

	if err != nil {
		http.Error(w, "Evénement innexistant", http.StatusInternalServerError)
	}

	comment.Event = int64(event_id)
	fmt.Println(comment.Event)

	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"error": "Entrées invalides",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if comment.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"error": "Contenu du commentaire vide",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	cookie, err := r.Cookie("jwt")

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jwtToken, err := auth.ValidateJWT(cookie.Value)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	rows, err := database.DB.Query(context.Background(), "SELECT id FROM users WHERE email=$1", jwtToken.Email)

	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println(err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		user_id := values[0]
		_, err = database.DB.Exec(context.Background(), "INSERT INTO comments (user_id, content, event_id) VALUES ($1, $2, $3)", user_id, comment.Content, comment.Event)

		if err != nil {
			log.Println(err)
			http.Error(w, "Database insert error", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]string{
		"success": "Commentaire ajouté",
	}
	json.NewEncoder(w).Encode(response)
}

func GetEventCommentsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	event_id := r.URL.Query().Get("event_id")

	rows, err := database.DB.Query(context.Background(), "SELECT comments.id, comments.content, users.name, comments.created_at FROM comments, users WHERE event_id =$1 and comments.user_id = users.id", event_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Erreur lors de la récupération des commentaires")
		return
	}

	defer rows.Close()

	var comments []models.CommentsJSON

	for rows.Next() {
		var comment models.CommentsJSON
		err := rows.Scan(&comment.ID, &comment.Content, &comment.User, &comment.Date)
		if err != nil {
			fmt.Printf("Erreur lors de la lecture des résultats %s \n", err)
			continue
		}
		comments = append(comments, comment)
	}

	response := map[string][]models.CommentsJSON{
		"comments": comments,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
