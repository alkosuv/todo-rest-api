package response

import "errors"

var (
	// ErrLoginUnavailable login unavailable
	ErrLoginUnavailable = errors.New("login unavailable")

	// ErrIncorrectLoginOrPassword incorrect login or password
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")

	// ErrIncorrectData incorrect data
	ErrIncorrectData = errors.New("incorrect data")

	// ErrNotAuthenticated not authenticated
	ErrNotAuthenticated = errors.New("not authenticated")

	// ErrSessions error session
	ErrSessions = errors.New("error session")
)
