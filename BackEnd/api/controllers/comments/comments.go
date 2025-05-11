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
	"strings"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	eventIDStr := r.URL.Query().Get("event_id")
	eventID, err := strconv.ParseInt(eventIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de l'événement invalide", http.StatusBadRequest)
		return
	}

	var comment models.Comments
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Entrées invalides", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(comment.Content) == "" {
		http.Error(w, "Contenu du commentaire vide", http.StatusBadRequest)
		return
	}
	comment.Event = eventID

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Token manquant",
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Format token invalide",
		})
		return
	}
	tokenStr := parts[1]

	jwtToken, err := auth.ValidateJWT(tokenStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Token invalide",
		})
		return
	}

	var userID int
	err = database.DB.QueryRow(context.Background(), "SELECT id FROM users WHERE email=$1", jwtToken.Email).Scan(&userID)
	if err != nil {
		log.Println("Erreur de récupération de l'utilisateur :", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Utilisateur inexistant",
		})
		return
	}

	_, err = database.DB.Exec(context.Background(), "INSERT INTO comments (user_id, content, event_id) VALUES ($1, $2, $3)", userID, comment.Content, comment.Event)
	if err != nil {
		log.Println("Erreur insertion commentaire :", err)
		http.Error(w, "Erreur lors de l'ajout du commentaire", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"success": "Commentaire ajouté",
	})
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
