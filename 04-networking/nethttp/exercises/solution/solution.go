// Package solution holds the reference solutions for the net/http exercises.
package solution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// RequireAPIKey rejects requests whose X-API-Key header is not key.
func RequireAPIKey(key string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != key {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if id < 1 || id > 100 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"id": id, "name": fmt.Sprintf("user%d", id)})
}

// NewAPI returns the application's HTTP handler:
//
//	GET /health      -> "ok" (no key needed)
//	GET /users/{id}  -> JSON user, requires the key "secret"
func NewAPI() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	mux.Handle("GET /users/{id}", RequireAPIKey("secret", http.HandlerFunc(getUser)))
	return mux
}
