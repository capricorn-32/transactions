package handler

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standard error response structure
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// SuccessResponse represents a standard success response structure
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// WriteErrorResponse writes a standardized error response
func WriteErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Error:   errorMessage,
	})
}

// WriteSuccessResponse writes a standardized success response
func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := SuccessResponse{
		Success: true,
	}

	if message != "" {
		response.Message = message
	}

	if data != nil {
		response.Data = data
	}

	json.NewEncoder(w).Encode(response)
}

// WriteBadRequestError is a convenience function for 400 Bad Request errors
func WriteBadRequestError(w http.ResponseWriter, errorMessage string) {
	WriteErrorResponse(w, http.StatusBadRequest, errorMessage)
}

// WriteNotFoundError is a convenience function for 404 Not Found errors
func WriteNotFoundError(w http.ResponseWriter, errorMessage string) {
	WriteErrorResponse(w, http.StatusNotFound, errorMessage)
}

// WriteCreatedResponse is a convenience function for 201 Created responses
func WriteCreatedResponse(w http.ResponseWriter, message string) {
	WriteSuccessResponse(w, http.StatusCreated, message, nil)
}

// WriteOKResponse is a convenience function for 200 OK responses with data
func WriteOKResponse(w http.ResponseWriter, data interface{}) {
	WriteSuccessResponse(w, http.StatusOK, "", data)
}
