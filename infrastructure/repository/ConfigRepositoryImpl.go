// infrastructure/repository/ConfigRepositoryImpl
package repository

import (
	"errors"
	"fmt"

	"github.com/Melbeck777/noq/internal/application/repository"
	"github.com/Melbeck777/noq/internal/application/validation"
	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/Melbeck777/noq/internal/domain/valueobject"
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
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return entity.Config{}, fmt.Errorf("%w: %v", repository.ErrConfigNotFound, err)
		}
		return entity.Config{}, fmt.Errorf("%w: %s\n", repository.ErrConfigParse, err)
	}

	var raw rawConfig
	err = v.Unmarshal(&raw)
	if err != nil {
		return entity.Config{}, fmt.Errorf("%w: unmarshall: %s\n", repository.ErrConfigInvalid, err)
	}

	// null check
	if err := validation.EmptyCheck(raw.MemoRoot, "memo_root", repository.ErrConfigInvalid); err != nil {
		return entity.Config{}, err
	}
	if err := validation.EmptyCheck(raw.ArticleRoot, "article_root", repository.ErrConfigInvalid); err != nil {
		return entity.Config{}, err
	}
	if err := validation.EmptyCheck(raw.Notion.Token, "notion.token", repository.ErrConfigInvalid); err != nil {
		return entity.Config{}, err
	}
	if err := validation.EmptyCheck(raw.Notion.DefaultMemoDB, "notion.default_memo_db", repository.ErrConfigInvalid); err != nil {
		return entity.Config{}, err
	}
	if len(raw.Notion.Databases) == 0 {
		return entity.Config{}, fmt.Errorf("%w: notion.databases is empty\n", repository.ErrConfigInvalid)
	}

	dbs, err := valueobject.NewNotionDatabases(raw.Notion.Databases)
	if err != nil {
		return entity.Config{}, fmt.Errorf("%w: notion.databases: %s\n", repository.ErrConfigInvalid, err)
	}

	defaultAlias, err := valueobject.NewDatabaseAlias(raw.Notion.DefaultMemoDB)
	if err != nil {
		return entity.Config{}, fmt.Errorf("%w: %s\n", repository.ErrConfigInvalid, err)
	}
	if _, ok := dbs.Get(defaultAlias); !ok {
		return entity.Config{}, fmt.Errorf("%w: default_memo_db not found in notion.databaes: %s\n", repository.ErrConfigInvalid, raw.Notion.DefaultMemoDB)
	}
	// TODO: ~/.config/noq/config.ymlがない時は作成する->オーケストレーション部分はusecaseで実装する
	// TODO: defaultMemoDBがない時に設定するように促す->別でこのファイル内で実装する
	// TODO: 単一項目のロードを実装する特にNotion.Databasesのみを取得する実装これはConfigよりも別の所だろうか？Configの値を読み出すからConfigに入れるのが適切なような気もしている
	cfg := entity.Config{
		MemoRoot:    raw.MemoRoot,
		ArticleRoot: raw.ArticleRoot,
		Notion: entity.NotionConfig{
			Token:         raw.Notion.Token,
			DefaultMemoDB: defaultAlias,
			Databases:     dbs,
		},
	}

	return cfg, nil
}
