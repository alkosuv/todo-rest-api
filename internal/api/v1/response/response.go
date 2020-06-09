package response

import (
	"encoding/json"
	"net/http"
)

// Response возвращает ответ от сервера
func Response(
	w http.ResponseWriter,
	statusCode int,
	data interface{},
) {
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Error возвращает ошибку от сервера
func Error(
	w http.ResponseWriter,
	statusCode int,
	err error,
) {
	Response(w, statusCode, map[string]string{"error": err.Error()})
}
