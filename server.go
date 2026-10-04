package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	widgets := []map[string]interface{}{
		{"id": 1, "name": "Widget A"},
		{"id": 2, "name": "Sprocket B"},
		{"id": 3, "name": "Gear C"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(widgets)

}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/api/widgets", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
