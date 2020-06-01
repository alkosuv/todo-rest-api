package middleware

import (
	"context"
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/ctxkey"

	"github.com/google/uuid"
)

// SetRequestID ...
func SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		// добавление uuid в заголовок
		w.Header().Set("X-Request-ID", id)

		req := r.WithContext(context.WithValue(
			r.Context(),
			ctxkey.CtxKeyRequestID,
			id,
		))
		next.ServeHTTP(w, req)
	})
}
