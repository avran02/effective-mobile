package dto

type UserDTO struct {
	PassportNumber string `json:"passportNumber"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}

type GetUserResponse struct {
	Users []UserDTO `json:"users"`
}

type CreateUserRequest struct {
	PassportNumber string `json:"passportNumber"`
}

type CreateUserResponse struct {
	ID int `json:"id"`
}
