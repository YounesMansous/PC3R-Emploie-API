package main

import (
	"api/controllers/auth"
	"api/controllers/comments"
	"api/controllers/events"
	"api/controllers/lines"
	"api/middlewares"
	"api/utils/database"
	"fmt"
	"net/http"
)

func loggingRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {

	database.ConnectDB()
	defer database.CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.LoginHandler)
	mux.HandleFunc("/register", auth.RegisterHandler)
	mux.HandleFunc("/lines/modes", lines.GetTransportModesHandler)
	mux.HandleFunc("/lines/modes/id", lines.GetTransportModeLinesIdsHandler)
	mux.HandleFunc("/events/line", events.GetLineEventsHandler)
	mux.HandleFunc("/events", events.GetEventHandler)
	mux.Handle("/comments/add", middlewares.JWTMiddleware(http.HandlerFunc(comments.AddCommentHandler)))
	mux.HandleFunc("/comments", comments.GetEventCommentsHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", loggingRequestMiddleware(mux))

}
