package main

import (
	"api/controllers/auth"
	"api/controllers/comments"
	"api/controllers/events"
	"api/controllers/lines"
	"api/middlewares"
	"api/utils"
	"api/utils/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/rs/cors"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	return os.Getenv(key)
}
func loggingRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {

	c := cron.New()

	prim_api_key := goDotEnvVariable("PRIM_API_KEY")
	c.AddFunc("@every 1h", func() {
		utils.PrimCall(prim_api_key)
	})

	c.AddFunc("@every 5s", func() {
		fmt.Println("API running... ", time.Now())
	})

	c.Start()

	databaseURL := goDotEnvVariable("DATABASE_URL")
	database.ConnectDB(databaseURL)
	defer database.CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.LoginHandler)
	mux.HandleFunc("/register", auth.RegisterHandler)
	mux.HandleFunc("/lines/modes", lines.GetTransportModesHandler)
	mux.HandleFunc("/lines/modes/id", lines.GetTransportModeLinesIdsHandler)
	mux.HandleFunc("/events/line", events.GetLineEventsHandler)
	mux.HandleFunc("/events", events.GetEventHandler)
	mux.HandleFunc("/logout", auth.LogoutHandler)
	mux.Handle("/comments/add", middlewares.JWTMiddleware(http.HandlerFunc(comments.AddCommentHandler)))
	mux.HandleFunc("/comments", comments.GetEventCommentsHandler)

	go func() {
		fmt.Println("Server started on port 8080")
		handler := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000", "https://traffik-two.vercel.app"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Authorization", "Content-Type"},
			AllowCredentials: true,
		}).Handler(loggingRequestMiddleware(mux))

		http.ListenAndServe(":8080", handler)

	}()

	select {}

}
