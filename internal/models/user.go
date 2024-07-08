package models

type User struct {
	ID             int
	PassportNumber string
	PassportSerie  string
	Name           string
	Surname        string
	Patronymic     string
	Address        string
}
