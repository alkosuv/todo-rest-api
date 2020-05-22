package response

import "errors"

var (
	// ErrLoginUnavailable ...
	ErrLoginUnavailable = errors.New("login unavailable")

	// ErrIncorrectLoginOrPassword ..
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")

	// ErrNotAuthenticated ...
	ErrNotAuthenticated = errors.New("not authenticated")
)
