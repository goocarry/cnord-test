package model

// User ...
type User struct {
	ID string `db:"user_id"`
	FirstName string `db:"firstname"`
	LastName string `db:"lastname"`
}