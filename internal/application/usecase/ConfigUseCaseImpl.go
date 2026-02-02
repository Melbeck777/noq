// internal/application/usecase/ConfigUseCaseImpl.go
package usecase

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Melbeck777/noq/internal/application/repository"
	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/Melbeck777/noq/internal/domain/valueobject"
)

// 最初にロードする
// 初期化をする
// 必要な情報が設定されていないのであれば設定を促すようにする
type ConfigUseCaseImpl struct {
	repo repository.ConfigRepository
}

func NewConfigUseCase(repo repository.ConfigRepository) *ConfigUseCaseImpl {
	return &ConfigUseCaseImpl{repo: repo}
}

func (u *ConfigUseCaseImpl) Load() (entity.Config, error) {
	cfg, err := u.repo.Load()
	if errors.Is(err, repository.ErrConfigParse) {
		return entity.Config{}, err
	} else if errors.Is(err, repository.ErrConfigNotFound) {
		return entity.Config{}, err
	}
	return cfg, nil
}

func (u *ConfigUseCaseImpl) Save(cfg entity.Config) error {
	return u.repo.Save(cfg)
}

func (u *ConfigUseCaseImpl) SaveOverWrite(cfg entity.Config) error {
	return u.repo.SaveOverWrite(cfg)
}

func (u *ConfigUseCaseImpl) DefaultValue() (entity.Config, error) {
	var cfg entity.Config
	home, err := os.UserHomeDir()
	if err != nil {
		return entity.Config{}, err
	}
	noq_dir := filepath.Join(home, "noq")
	cfg.MemoRoot = filepath.Join(noq_dir, "memo")
	cfg.ArticleRoot = filepath.Join(noq_dir, "article")
	alias, err := valueobject.NewDatabaseAlias("None")
	if err != nil {
		return entity.Config{}, err
	}
	cfg.Notion.DefaultMemoDB = *alias
	return cfg, nil
}
