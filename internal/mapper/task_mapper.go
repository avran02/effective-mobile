package mapper

import (
	"github.com/avran02/effective-mobile/internal/dto"
	"github.com/avran02/effective-mobile/internal/models"
)

func FromTaskDTOToModel(dto dto.CreateTaskRequest) models.Task {
	return models.Task{
		UserID:      dto.UserID,
		Name:        dto.Name,
		Description: dto.Description,
	}
}
