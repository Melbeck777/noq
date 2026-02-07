package mapper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/Melbeck777/noq/internal/presentation/dto"
	"github.com/Melbeck777/noq/internal/presentation/mapper"
)

func TestInitRequestDtoToConfig_whenValidDto_thenMappedConfig(t *testing.T) {
	// Given
	m := mapper.NewConfigMapper()

	in, err := dto.NewInitRequestDto()
	if err != nil {
		t.Fatalf("NewInitRequestDto() error = %v", err)
	}
	in.MemoRoot = "/tmp/noq/memo"
	in.ArticleRoot = "/tmp/noq/article"
	in.NotionToken = "secret"
	in.DefaultMemoDB = "memo"
	in.Databases = map[string]string{
		"memo": "dbid_1",
		"sub":  "dbid_2",
	}

	// When
	cfg, err := m.InitRequestDtoToConfig(in)
	// Then
	if err != nil {
		t.Fatalf("InitRequestDtoToConfig() error = %v", err)
	}
	if cfg.MemoRoot != in.MemoRoot {
		t.Errorf("MemoRoot = %q, want %q", cfg.MemoRoot, in.MemoRoot)
	}
	if cfg.ArticleRoot != in.ArticleRoot {
		t.Errorf("ArticleRoot = %q, want %q", cfg.ArticleRoot, in.ArticleRoot)
	}

	if cfg.Notion.Token != in.NotionToken {
		t.Errorf("Notion.Token = %q, want %q", cfg.Notion.Token, in.NotionToken)
	}
	if got := cfg.Notion.DefaultMemoDB.Value(); got != in.DefaultMemoDB {
		t.Errorf("Notion.DefaultMemoDB = %q, want %q", got, in.DefaultMemoDB)
	}
	if got := cfg.Notion.Databases.GetDBMap(); !reflect.DeepEqual(got, in.Databases) {
		t.Errorf("Notion.Databases map = %#v, want %#v", got, in.Databases)
	}
	fmt.Println(in)
	fmt.Println(cfg)
}

func TestInitRequestDtoToConfig_whenInvalidDefaultMemoDB_thenReturnError(t *testing.T) {
	// Given
	m := mapper.NewConfigMapper()

	in, err := dto.NewInitRequestDto()
	if err != nil {
		t.Fatalf("NewInitRequestDto() error = %v", err)
	}
	in.MemoRoot = "/tmp/noq/memo"
	in.ArticleRoot = "/tmp/noq/article"
	in.NotionToken = "secret"
	in.DefaultMemoDB = "" // valueobject.NewDatabaseAlias が弾く想定
	in.Databases = map[string]string{"memo": "dbid_1"}

	// When
	_, err = m.InitRequestDtoToConfig(in)

	// Then
	if err == nil {
		t.Fatalf("InitRequestDtoToConfig() expected error, got nil")
	}
}

func TestEntityConfigToInitRequestDto_whenValidConfig_thenMappedDto(t *testing.T) {
	// Given
	m := mapper.NewConfigMapper()

	notion, err := entity.NewNotionConfig(
		"secret",
		"memo",
		map[string]string{"memo": "dbid_1"},
	)
	if err != nil {
		t.Fatalf("NewNotionConfig() error = %v", err)
	}

	cfg := entity.Config{
		MemoRoot:    "/tmp/noq/memo",
		ArticleRoot: "/tmp/noq/article",
		Notion:      notion,
	}

	// When
	out, err := m.EntityConfigToInitRequestDto(cfg)
	// Then
	if err != nil {
		t.Fatalf("EntityConfigToInitRequestDto() error = %v", err)
	}
	if out.MemoRoot != cfg.MemoRoot {
		t.Errorf("MemoRoot = %q, want %q", out.MemoRoot, cfg.MemoRoot)
	}
	if out.ArticleRoot != cfg.ArticleRoot {
		t.Errorf("ArticleRoot = %q, want %q", out.ArticleRoot, cfg.ArticleRoot)
	}
	if out.DefaultMemoDB != cfg.Notion.DefaultMemoDB.Value() {
		t.Errorf("DefaultMemoDB = %q, want %q", out.DefaultMemoDB, cfg.Notion.DefaultMemoDB.Value())
	}
	if !reflect.DeepEqual(out.Databases, cfg.Notion.Databases.GetDBMap()) {
		t.Errorf("Databases = %#v, want %#v", out.Databases, cfg.Notion.Databases.GetDBMap())
	}
}

func TestEntityConfigToInitRequestDto_whenDatabasesNil_thenReturnEmptyMap(t *testing.T) {
	// Given
	m := mapper.NewConfigMapper()

	// NotionDatabases.GetDBMap() が nil を返しうる前提のテスト。
	// もし実装上 nil を返さないなら、このテストは削除でOK。
	notion, err := entity.NewNotionConfig("secret", "memo", nil)
	if err != nil {
		t.Fatalf("NewNotionConfig() error = %v", err)
	}

	cfg := entity.Config{
		MemoRoot:    "/tmp/noq/memo",
		ArticleRoot: "/tmp/noq/article",
		Notion:      notion,
	}

	// When
	out, err := m.EntityConfigToInitRequestDto(cfg)
	// Then
	if err != nil {
		t.Fatalf("EntityConfigToInitRequestDto() error = %v", err)
	}
	if out.Databases == nil {
		t.Fatalf("Databases should not be nil")
	}
	if len(out.Databases) != 0 {
		t.Fatalf("Databases len = %d, want 0", len(out.Databases))
	}
}
