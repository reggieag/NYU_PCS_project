package openapi

import (
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(authUrl string) func(http.Handler) http.Handler {
	log.Printf("initializing auth middleware using url: %s", authUrl)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("authorizing request")
			token := getBearerToken(r)
			if token == "" {
				http.Error(w, "Token not found", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})

	}
}

func getBearerToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	splitToken := strings.Split(bearer, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return splitToken[1]
}
