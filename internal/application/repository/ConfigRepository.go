package repository

import "github.com/Melbeck777/noq/internal/domain/entity"

type ConfigRepository interface {
	Load() (entity.Config, error)
}
