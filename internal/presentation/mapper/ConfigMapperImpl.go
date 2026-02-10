package mapper

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/Melbeck777/noq/internal/presentation/dto"
)

type ConfigMapperImpl struct{}

func NewConfigMapper() *ConfigMapperImpl {
	return &ConfigMapperImpl{}
}

func (m *ConfigMapperImpl) InitRequestDtoToConfig(dto dto.InitRequestDto) (entity.Config, error) {
	var err error
	notionConfig, err := entity.NewNotionConfig(dto.NotionToken, dto.DefaultMemoDB, dto.Databases)
	if err != nil {
		return entity.Config{}, err
	}
	res := entity.Config{dto.MemoRoot, dto.ArticleRoot, notionConfig}
	fmt.Println(res)

	return res, nil
}

func (m *ConfigMapperImpl) EntityConfigToInitRequestDto(cfg entity.Config) (dto.InitRequestDto, error) {
	if m == nil {
		panic("ConfigMpapperImpl is nil")
	}
	res, err := dto.NewInitRequestDto()
	if err != nil {
		return dto.InitRequestDto{}, err
	}
	res.MemoRoot = cfg.MemoRoot
	res.ArticleRoot = cfg.ArticleRoot
	if cfg.Notion.DefaultMemoDB != nil {
		res.DefaultMemoDB = cfg.Notion.DefaultMemoDB.Value()
	}
	res.Databases = cfg.Notion.Databases.GetDBMap()
	if res.Databases == nil {
		res.Databases = make(map[string]string)
	}
	return res, nil
}

func prettyPrint(v any) {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf.String())
}
