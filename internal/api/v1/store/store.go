package store

// Store ...
type Store interface {
	Todo() TodoRepository
	User() UserRepository
}
