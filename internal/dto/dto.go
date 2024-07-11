package dto

type UserDTO struct {
	ID             int    `json:"id"`
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

type CreateTaskRequest struct {
	UserID      int    `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	ID int `json:"id"`
}
