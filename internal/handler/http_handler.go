package handler

import (
	"encoding/json"
	"net/http"
	"short-url-service/internal/storage"
	"strings"
)

func PostHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method, expected POST", http.StatusMethodNotAllowed)
			return
		}

		var req map[string]string
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		url := req["url"]
		if url == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		shortURL, err := storage.Post(url)
		if err != nil {
			http.Error(w, "Error saving URL", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"short_url": shortURL,
		})
	}
}

func GetHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method, expected Get", http.StatusMethodNotAllowed)
			return
		}

		shortURL := strings.TrimPrefix(r.URL.Path, "/api/url/")
		if shortURL == "" {
			http.Error(w, "ShortURL is required", http.StatusBadRequest)
			return
		}

		originalURL, err := storage.Get(shortURL)
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"url": originalURL,
		})
	}
}
