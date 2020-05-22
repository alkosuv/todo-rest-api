package response

import (
	"encoding/json"
	"net/http"
)

// Response ...
func Response(
	w http.ResponseWriter,
	r *http.Request,
	statusCode int,
	data interface{},
) {
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Error ...
func Error(
	w http.ResponseWriter,
	r *http.Request,
	statusCode int,
	err error,
) {
	Response(w, r, statusCode, map[string]string{"error": err.Error()})
}
