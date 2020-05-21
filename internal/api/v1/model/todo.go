package model

// Todo ...
type Todo struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id,omitempty"`
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	DateCreate string `json:"date_create"`
}
