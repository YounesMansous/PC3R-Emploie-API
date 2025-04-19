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
	fmt.Printf("endpoint récupérant les ids des lignes du mode de transport")
}
