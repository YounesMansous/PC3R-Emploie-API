package auth

import (
	"api/models"
	"api/utils"
	"api/utils/database"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(utils.GoDotEnvVariable("JWT_KEY"))

type JwtPayload struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type LoginData struct {
	Email    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user models.Users
	var loginData LoginData

	err := json.NewDecoder(r.Body).Decode(&loginData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("%s \n", err)
		return
	}

	rows, err := database.DB.Query(context.Background(), "SELECT email, password FROM users where email=$1", loginData.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("%s \n", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Email, &user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("%s \n", err)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{
			"error": "Utilisateur introuvable",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{
			"error": "Mot de passe incorrect",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	token, err := GenerateJWT(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("%s \n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"success": "Utilisateur connecté",
		"token":   token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user models.Users

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hashed_password, err := HashPassword(user.Password)

	if err != nil {
		fmt.Printf("Json error %s", err)
		return
	}

	user.Password = hashed_password

	_, err = database.DB.Exec(context.Background(), "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)

	if err != nil {
		log.Println(err)
		http.Error(w, "Database insert error", http.StatusInternalServerError)
		return
	}

	token, err := GenerateJWT(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("%s\n", err)
		return
	}

	response := map[string]string{
		"success": "Compte enregistré",
		"token":   token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Token d'autorisation manquant",
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Format de token invalide",
		})
		return
	}

	token := parts[1]
	_, err := ValidateJWT(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Token invalide",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Déconnexion réussie",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JwtPayload{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*JwtPayload, error) {
	claims := &JwtPayload{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
