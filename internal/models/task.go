package models

type Task struct {
	ID          int
	Name        string
	Description string

	UserID int
}
