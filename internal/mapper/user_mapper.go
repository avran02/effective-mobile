package mapper

import (
	"github.com/avran02/effective-mobile/internal/dto"
	"github.com/avran02/effective-mobile/internal/models"
	"github.com/avran02/effective-mobile/utils"
)

func FromUserModelToUserDTO(model models.User) dto.UserDTO {
	hiddenPassport := utils.HidePassportNumber(model.PassportNumber)

	return dto.UserDTO{
		ID:             model.ID,
		PassportNumber: hiddenPassport,
		Name:           model.Name,
		Surname:        model.Surname,
		Patronymic:     model.Patronymic,
		Address:        model.Address,
	}
}

func FromUserDTOToUserModel(dto dto.UserDTO) models.User {
	return models.User{
		ID:             dto.ID,
		PassportNumber: dto.PassportNumber,
		Name:           dto.Name,
		Surname:        dto.Surname,
		Patronymic:     dto.Patronymic,
		Address:        dto.Address,
	}
}
