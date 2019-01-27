package main

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/pw-test/models"
	"net/http"
	"strings"
)

// NOT SAFE. IT'S JUST FOR TEST TASK !
const TOKEN = "thisIsTheJwtSecretPassword"

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/sign-in", "/api/sign-up"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{Error: "Missing auth token"})
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{Error: "Invalid/Malformed auth token"})
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(TOKEN), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{Error: "Malformed authentication token"})
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Response{Error: "Token is not valid"})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.Email)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func getToken(email string) string {
	tk := &models.Token{Email: email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(TOKEN))
	return tokenString
}
