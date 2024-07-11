package mapper

import "github.com/avran02/effective-mobile/internal/models"

func ToDatabaseUpdateUserData(model models.User) map[string]string {
	res := make(map[string]string)

	if model.PassportNumber != "" {
		res["passport_number"] = model.PassportNumber
	}

	if model.Name != "" {
		res["name"] = model.Name
	}

	if model.Surname != "" {
		res["surname"] = model.Surname
	}

	if model.Patronymic != "" {
		res["patronymic"] = model.Patronymic
	}

	if model.Address != "" {
		res["address"] = model.Address
	}

	return res
}
