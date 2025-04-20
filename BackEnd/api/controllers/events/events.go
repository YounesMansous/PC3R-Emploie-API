package events

import (
	"api/models"
	"api/utils/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetLineEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var events []models.EventsJSON

	line_id := r.URL.Query().Get("id_line")

	rows, err := database.DB.Query(context.Background(), "SELECT events.id, events.title, events.message, events.publication_time FROM events, lines WHERE lines.line_id =$1 AND lines.line_id = events.line_id ORDER BY events.publication_time DESC", line_id)

	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var event models.EventsJSON
		err = rows.Scan(&event.ID, &event.Titre, &event.Message, &event.Date)
		if err != nil {
			fmt.Printf("Erreur lors de la lecture des résultats %s \n", err)
			continue
		}
		events = append(events, event)
	}

	response := map[string][]models.EventsJSON{
		"events": events,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func GetEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	event_id := r.URL.Query().Get("id")

	rows, err := database.DB.Query(context.Background(), "SELECT id, events.title, events.message, events.publication_time FROM events  WHERE id =$1", event_id)

	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var event models.EventsJSON

	if rows.Next() {
		err = rows.Scan(&event.ID, &event.Titre, &event.Message, &event.Date)
		if err != nil {
			fmt.Printf("Erreur lors de la lecture des résultats %s \n", err)
		}
		response := map[string]models.EventsJSON{
			"event": event,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]*models.EventsJSON{
		"event": nil,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
