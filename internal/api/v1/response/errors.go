package response

import "errors"

var (
	// ErrLoginUnavailable ...
	ErrLoginUnavailable = errors.New("login unavailable")

	// ErrIncorrectLoginOrPassword ..
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")

	// ErrIncorrectData ...
	ErrIncorrectData = errors.New("incorrect data")

	// ErrNotAuthenticated ...
	ErrNotAuthenticated = errors.New("not authenticated")

	// ErrSessions ...
	ErrSessions = errors.New("error session")
)
