package httpx

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Success bool      `json:"success"`
	Data    any       `json:"data"`
	Error   *ApiError `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ApiResponse{
		Success: true,
		Data:    data,
	})
}

func Fail(w http.ResponseWriter, e *ApiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	_ = json.NewEncoder(w).Encode(ApiResponse{
		Success: false,
		Error:   e,
	})
}
