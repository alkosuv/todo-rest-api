package store

// Store интерфейс описывающий функции запросов к БД
type Store interface {
	Todo() TodoRepository
	User() UserRepository
}
