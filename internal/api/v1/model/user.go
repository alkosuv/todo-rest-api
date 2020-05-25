package model

// User ...
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
}

// IsNil ...
func (u *User) IsNil() bool {
	return *u == User{}
}
