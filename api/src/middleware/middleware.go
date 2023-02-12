package middleware

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
)

// Logs the requests
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s ", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
		}
		nextFunc(w, r)
	}
}
