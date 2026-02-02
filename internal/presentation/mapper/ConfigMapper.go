package mapper

import (
	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/Melbeck777/noq/internal/presentation/dto"
)

type ConfigMapper interface {
	InitRequestDtoToEntityConfig(dto.InitRequestDto) (entity.Config, error)
	EntityConfigToInitRequestDto(entity.Config) (dto.InitRequestDto, error)
}
