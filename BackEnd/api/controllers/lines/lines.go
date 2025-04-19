package lines

import (
	"api/models"
	"api/utils/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetTransportModesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var modes []models.TypeTransportsJSON

	rows, err := database.DB.Query(context.Background(), "SELECT DISTINCT type FROM lines ORDER BY type desc")

	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var mode models.TypeTransportsJSON
		err := rows.Scan(&mode.TYPE)
		if err != nil {
			fmt.Printf("Erreur lors de la lecture des résultats %s \n", err)
			continue
		}
		modes = append(modes, mode)
	}

	response := map[string][]models.TypeTransportsJSON{
		"modes": modes,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetTransportModeLinesIdsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mode := r.URL.Query().Get("mode")

	rows, err := database.DB.Query(context.Background(), "SELECT line_id, line_name FROM lines WHERE type =$1 ORDER BY line_name ASC", mode)

	if err != nil {
		log.Println(err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var lines []models.LineIDsJSON

	for rows.Next() {
		var line models.LineIDsJSON
		rows.Scan(&line.ID, &line.NAME)
		lines = append(lines, line)
	}

	response := map[string][]models.LineIDsJSON{
		"lines": lines,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
