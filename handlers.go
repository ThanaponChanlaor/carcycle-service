package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("write response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"message": message})
}

func (s *Store) handleListCars(w http.ResponseWriter, r *http.Request) {
	cars, err := s.GetAllCars()
	if err != nil {
		log.Printf("list cars: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to load cars")
		return
	}
	writeJSON(w, http.StatusOK, cars)
}

func (s *Store) handleGetCar(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	car, err := s.GetCarByID(id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, fmt.Sprintf("car %q was not found", id))
		return
	}
	if err != nil {
		log.Printf("get car %s: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to load car")
		return
	}
	writeJSON(w, http.StatusOK, car)
}

func (s *Store) handleGetHistory(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if _, err := s.GetCarByID(id); errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, fmt.Sprintf("car %q was not found", id))
		return
	}

	entries, err := s.GetHistory(id)
	if err != nil {
		log.Printf("get history %s: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to load history")
		return
	}
	writeJSON(w, http.StatusOK, entries)
}

func (s *Store) handleUpdateCar(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var input CarUpdateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	updated, err := s.UpdateCar(id, input)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, fmt.Sprintf("car %q was not found", id))
		return
	}
	if err != nil {
		log.Printf("update car %s: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to update car")
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "service": "carcycle-service"})
}
