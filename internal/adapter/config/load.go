package config

import (
	"fmt"
	"os"

	"github.com/Melbeck777/noq/internal/usecase"

	"github.com/spf13/viper"
)

type NotionConfig struct {
	DefaultMemoDB string            `mapstructure:"default_memo_db" yaml:"default_memo_db"`
	Databases     map[string]string `mapstructure:"databases" yaml:"databases"`
}

type fileConfig struct {
	MemoRoot    string       `mapstructure:"memo_root" yaml:"memo_root"`
	ArticleRoot string       `mapstructure:"article_root" yaml:"article_root"`
	Notion      NotionConfig `mapstructure:"notion" yaml:"notion"`
}

func (c *fileConfig) MemoRootPath() string                  { return c.MemoRoot }
func (c *fileConfig) ArticleRootPath() string               { return c.ArticleRoot }
func (c *fileConfig) NotionMemoDBId() string                { return c.Notion.DefaultMemoDB }
func (c *fileConfig) NotionMemoDBListId() map[string]string { return c.Notion.Databases }

func Load(path string) (usecase.Config, error) {
	v := viper.New()

	if _, exist_check := os.Stat(path); exist_check == nil {
		return nil, fmt.Errorf("failed to %s is not exist:", path)
	}

	var raw fileConfig
	if err := v.Unmarshal(&raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validate(&raw); err != nil {
		return nil, err
	}

	return &raw, nil
}

func validate(cfg *fileConfig) error {
	var missing []string

	if cfg.MemoRootPath() == "" {
		missing = append(missing, "memo_root")
	}

	if cfg.ArticleRootPath() == "" {
		missing = append(missing, "article_root")
	}
	if cfg.NotionMemoDBId() == "" {
		missing = append(missing, "ntoion.DefaultMemoDB")
	}

	if len(cfg.NotionMemoDBListId()) == 0 {
		missing = append(missing, "notion.Databases")
	}

	if len(missing) > 0 {
		return fmt.Errorf("config: missing required fields: %v", missing)
	}

	return nil
}
