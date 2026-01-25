package repository

import (
	"fmt"

	"github.com/Melbeck777/noq/internal/application/repository"
	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/spf13/viper"
)

type ConfigRepositoryImpl struct {
	configPath string
}

func NewConfigRepositoryImpl(configPath string) repository.ConfigRepository {
	return &ConfigRepositoryImpl{configPath: configPath}
}

type rawConfig struct {
	MemoRoot    string `mapstructure:"memo_root"`
	ArticleRoot string `mapstructure:"article_root"`
	Notion      struct {
		Token         string            `mapstructure:"token"`
		DefaultMemoDB string            `mapstructure:"default_memo_db"`
		Databases     map[string]string `mapstructure:"databases"`
	} `mapstructure:"notion"`
}

func (r *ConfigRepositoryImpl) Load() (entity.Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME/.config/noq")
	err := v.ReadInConfig()
	if err != nil {
		return entity.Config{}, fmt.Errorf("Setting file read erro: %s\n", err)
	}

	var cfg entity.Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		return entity.Config{}, fmt.Errorf("unmarshall erro: %s\n", err)
	}
	// config.ymlから値を取り出す
	// 値の検証を行う
	// configにマッピングする

	cfg := entity.Config{}
	return cfg, nil
}
