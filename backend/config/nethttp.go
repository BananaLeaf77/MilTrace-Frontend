package config

import (
	"encoding/json"
	"net/http"
)

// Helper functions for error responses
func NetHTTPBadRequest(w http.ResponseWriter, message string, err string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "false",
		"message": message,
		"error":   err,
		"status":  400,
	})
}

func NetHTTPInternalServerError(w http.ResponseWriter, message string, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "false",
		"message": message,
		"error":   err,
		"status":  500,
	})
}

func NetHTTPStatusOK(w http.ResponseWriter, message string, data *any) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "true",
		"message": message,
		"status":  200,
		"data":    data,
	})
}

func NetHTTPStatusCreated(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "true",
		"message": message,
		"status":  201,
	})
}

func NetHTTPStatusNoContent(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "true",
		"message": message,
		"status":  204,
		"data":    nil,
	})
}

func NetHTTPUnauthorized(w http.ResponseWriter, message string, err string) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "false",
		"message": message,
		"error":   err,
		"status":  401,
	})
}

func NetHTTPForbidden(w http.ResponseWriter, message string, err string) {
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "false",
		"message": message,
		"error":   err,
		"status":  403,
	})
}

func NetHTTPNotFound(w http.ResponseWriter, message string, err string) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]any{
		"success": "false",
		"message": message,
		"error":   err,
		"status":  404,
	})
}
