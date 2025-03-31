package middlewares

import (
	"api/controllers/auth"
	"net/http"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Le token est dans le format "Bearer <token>", on récupère la seconde partie
		token := cookie.Value

		// Valider le token
		_, err = auth.ValidateJWT(token)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Continuer vers le handler
		next.ServeHTTP(w, r)
	})
}
