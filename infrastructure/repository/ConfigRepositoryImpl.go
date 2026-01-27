// infrastructure/repository/ConfigRepositoryImpl
package repository

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Melbeck777/noq/internal/application/repository"
	"github.com/Melbeck777/noq/internal/application/validation"
	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/Melbeck777/noq/internal/domain/valueobject"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type ConfigRepositoryImpl struct {
	configPath string
}

func NewConfigRepositoryImpl(configPath string) repository.ConfigRepository {
	return &ConfigRepositoryImpl{configPath: configPath}
}

type rawConfig struct {
	MemoRoot    string `yaml:"memo_root" mapstructure:"memo_root"`
	ArticleRoot string `yaml:"article_root" mapstructure:"article_root"`
	Notion      struct {
		Token         string            `yaml:"token" mapstructure:"token"`
		DefaultMemoDB string            `yaml:"default_memo_db" mapstructure:"default_memo_db"`
		Databases     map[string]string `yaml:"databases" mapstructure:"databases"`
	} `yaml:"notion" mapstructure:"notion"`
}

func (r *ConfigRepositoryImpl) Load() (entity.Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetConfigFile(r.configPath)
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

func (r *ConfigRepositoryImpl) configEntityToRawConfig(cfg entity.Config) rawConfig {
	raw := rawConfig{}
	raw.MemoRoot = cfg.MemoRoot
	raw.ArticleRoot = cfg.ArticleRoot
	raw.Notion.Token = cfg.Notion.Token
	raw.Notion.DefaultMemoDB = string(cfg.Notion.DefaultMemoDB)
	raw.Notion.Databases = map[string]string{}
	for k, v := range cfg.Notion.Databases.GetDBMap() {
		raw.Notion.Databases[string(k)] = string(v)
	}
	return raw
}

func (r *ConfigRepositoryImpl) Save(cfg entity.Config) error {
	return r.save(cfg, false)
}

func (r *ConfigRepositoryImpl) SaveOverwrite(cfg entity.Config) error {
	return r.save(cfg, true)
}

func (r *ConfigRepositoryImpl) save(cfg entity.Config, overwrite bool) error {
	raw := r.configEntityToRawConfig(cfg)
	b, err := yaml.Marshal(raw)
	if err != nil {
		return err
	}
	if !overwrite {
		if _, err := os.Stat(r.configPath); err == nil {
			return fmt.Errorf("config.yml is already exist")
		}
	}
	fmt.Printf(raw.MemoRoot)
	if err := ioutil.WriteFile(r.configPath, b, 0o600); err != nil {
		return err
	}
	return nil
}
