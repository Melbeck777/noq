// internal/application/usecase/ConfigUseCase.go
package usecase

import (
	"errors"

	"github.com/Melbeck777/noq/internal/application/repository"
	"github.com/Melbeck777/noq/internal/domain/entity"
)

// 最初にロードする
// 初期化をする
// 必要な情報が設定されていないのであれば設定を促すようにする
type ConfigUseCase struct {
	repo repository.ConfigRepository
}

func NewConfigUseCase(repo repository.ConfigRepository) *ConfigUseCase {
	return &ConfigUseCase{repo: repo}
}

func (u *ConfigUseCase) Load() (entity.Config, error) {
	cfg, err := u.repo.Load()
	if errors.Is(err, repository.ErrConfigParse) {
		return entity.Config{}, err
	} else if errors.Is(err, repository.ErrConfigNotFound) {
		return entity.Config{}, err
	} else if errors.Is(err, repository.ErrConfigParse) {
		return entity.Config{}, err
	}
	return cfg, nil
}
