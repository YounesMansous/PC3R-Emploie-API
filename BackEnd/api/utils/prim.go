package utils

import (
	"api/utils/database"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Pagination struct {
	TotalResult  int `json:"total_result"`
	StartPage    int `json:"start_page"`
	ItemsPerPage int `json:"items_per_page"`
	ItemsOnPage  int `json:"items_on_page"`
}

type Disruption struct {
	UpdatedAt string `json:"updated_at"`
	Messages  []struct {
		Text    string `json:"text"`
		Channel struct {
			Name string `json:"name"`
		} `json:"channel"`
	} `json:"messages"`
	ImpactedObjects []struct {
		PtObject struct {
			ID   string `json:"id"`
			LINE struct {
				ID string `json:"id"`
			} `json:"line"`
		} `json:"pt_object"`
	} `json:"impacted_objects"`
}

type APIResponse struct {
	Pagination  Pagination   `json:"pagination"`
	Disruptions []Disruption `json:"disruptions"`
}

func PrimCall(apiKey string) {

	fmt.Println("Mise à jour des évenements...")
	page := 0
	total := 1

	_, err := database.DB.Exec(context.Background(),
		"DELETE  FROM events",
	)

	if err != nil {
		fmt.Println("Erreur nettoyage table", err)
	}

	for page*10 < total {
		url := `https://prim.iledefrance-mobilites.fr/marketplace/v2/navitia/line_reports/line_reports?depth=3&count=100&start_page=` + strconv.Itoa(page) + `&disable_geojson=false&tags%5B%5D=Actualit%C3%A9`
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Erreur de requête:", err)
			continue
		}
		req.Header.Add("apikey", apiKey)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Erreur d'appel HTTP:", err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		var data APIResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Erreur JSON:", err)
			continue
		}

		total = data.Pagination.TotalResult

		// Traitement des disruptions
		for _, disruption := range data.Disruptions {
			updatedAt := disruption.UpdatedAt
			titre := ""
			message := ""
			for _, msg := range disruption.Messages {
				switch msg.Channel.Name {
				case "titre":
					titre = msg.Text
				case "moteur":
					message = msg.Text
				}
			}

			for _, obj := range disruption.ImpactedObjects {
				lineID := strings.TrimPrefix(obj.PtObject.LINE.ID, "line:")

				_, err := database.DB.Exec(context.Background(),
					"INSERT INTO events (line_id, title, message, publication_time) VALUES ($1, $2, $3, $4)",
					lineID, titre, message, updatedAt,
				)
				if err != nil {
					fmt.Println("Erreur d'insertion pour la ligne", lineID, ":", err)
				}
			}

		}

		page++
	}
	fmt.Println("Tous les résultats ont été insérés.")
}
