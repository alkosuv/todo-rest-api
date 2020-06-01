package log

import (
	"net/http"

	"github.com/gen95mis/todo-rest-api/internal/api/v1/ctxkey"
	"github.com/sirupsen/logrus"
)

// Error ...
func Error(logger *logrus.Logger, r *http.Request, code int, err error) {
	logger.WithFields(logrus.Fields{
		"uuid":   r.Context().Value(ctxkey.CtxKeyRequestID),
		"url":    r.RequestURI,
		"method": r.Method,
		"code":   code,
		"error":  err.Error(),
	}).Error()
}

// Info ...
func Info(logger *logrus.Logger, r *http.Request, code int, message interface{}) {
	logger.WithFields(logrus.Fields{
		"uuid":   r.Context().Value(ctxkey.CtxKeyRequestID),
		"url":    r.RequestURI,
		"method": r.Method,
		"code":   code,
		"msg":    message,
	}).Info()
}
