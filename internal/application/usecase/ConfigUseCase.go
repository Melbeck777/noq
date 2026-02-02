package usecase

import "github.com/Melbeck777/noq/internal/domain/entity"

type ConfigUseCase interface {
	Load() (entity.Config, error)
	Save(entity.Config) error
	SaveOverWrite(entity.Config) error
	DefaultValue() (entity.Config, error)
}
