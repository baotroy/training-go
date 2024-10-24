package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func resonseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println(message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, errorResponse{
		Error: message,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to marshal json %v", err)
		w.WriteHeader((http.StatusInternalServerError))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
