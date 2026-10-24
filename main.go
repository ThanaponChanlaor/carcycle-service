package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	loadDotEnv(".env")

	port := getEnv("PORT", "8080")
	dbPath := getEnv("DB_PATH", "carcycle.db")

	db := mustOpenDB(dbPath)
	defer db.Close()

	store := &Store{db: db}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleHealth)
	mux.HandleFunc("GET /api/cars", store.handleListCars)
	mux.HandleFunc("GET /api/cars/{id}", store.handleGetCar)
	mux.HandleFunc("GET /api/cars/{id}/history", store.handleGetHistory)
	mux.HandleFunc("PATCH /api/cars/{id}", store.handleUpdateCar)

	addr := ":" + port
	log.Printf("carcycle-service listening on %s (db: %s)", addr, dbPath)
	log.Fatal(http.ListenAndServe(addr, withCORS(withLogging(mux))))
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

// loadDotEnv is a minimal .env reader (no external dependency): it sets
// process env vars from KEY=VALUE lines, skipping ones already set in the
// real environment. Missing file is not an error — real env vars/defaults
// still apply.
func loadDotEnv(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}
}
