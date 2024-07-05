package mapper

import (
	"github.com/avran02/effective-mobile/internal/dto"
	"github.com/avran02/effective-mobile/internal/models"
	"github.com/avran02/effective-mobile/utils"
)

func FromUserModelToUserDTO(model models.User) dto.UserDTO {
	fullPassport := model.PassportSerie + model.PassportNumber
	hiddenPassport := utils.HidePassportNumber(fullPassport)

	return dto.UserDTO{
		PassportNumber: hiddenPassport,
		Name:           model.Name,
		Surname:        model.Surname,
		Patronymic:     model.Patronymic,
		Address:        model.Address,
	}
}

func FromUserDTOToUserModel(dto dto.UserDTO) models.User {
	return models.User{
		PassportSerie:  dto.PassportNumber[:4],
		PassportNumber: dto.PassportNumber[4:],
		Name:           dto.Name,
		Surname:        dto.Surname,
		Patronymic:     dto.Patronymic,
		Address:        dto.Address,
	}
}
