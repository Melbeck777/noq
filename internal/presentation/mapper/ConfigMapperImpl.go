package mapper

import (
	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/Melbeck777/noq/internal/domain/valueobject"
	"github.com/Melbeck777/noq/internal/presentation/dto"
)

type ConfigMapperImpl struct{}

func NewConfigMapper() *ConfigMapperImpl {
	return &ConfigMapperImpl{}
}

func (m *ConfigMapperImpl) InitRequestDtoToConfig(dto dto.InitRequestDto) (entity.Config, error) {
	var res entity.Config
	var err error
	res.MemoRoot = dto.MemoRoot
	res.ArticleRoot = dto.ArticleRoot
	memo_db, err := valueobject.NewDatabaseAlias(dto.DefaultMemoDB)
	if err != nil {
		return entity.Config{}, err
	}
	res.Notion.DefaultMemoDB = *memo_db
	res.Notion.Token = dto.NotionToken
	res.Notion.Databases, err = valueobject.NewNotionDatabases(dto.Databases)
	if err != nil {
		return entity.Config{}, err
	}
	return res, nil
}

func (m *ConfigMapperImpl) EntityConfigToInitRequestDto(cfg entity.Config) (dto.InitRequestDto, error) {
	var res dto.InitRequestDto
	res.MemoRoot = cfg.MemoRoot
	res.ArticleRoot = cfg.ArticleRoot
	res.DefaultMemoDB = cfg.Notion.DefaultMemoDB.Value()
	return res, nil
}
