package middlewares

import (
	"errors"
	"github.com/MD-ARMAN-Shanto/gostack/api/auth"
	"github.com/MD-ARMAN-Shanto/gostack/api/responses"
	"net/http"
)

func SetMiddlewareJson(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		next(w, r)
	}
}
