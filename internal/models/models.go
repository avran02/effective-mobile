package models

type User struct {
	ID             int
	PassportSerie  string
	PassportNumber string
	Name           string
	Surname        string
	Patronymic     string
	Address        string
}

type Task struct {
	TaskID   int
	UserID   int
	TaskName string
	Hours    int
	Minutes  int
}
