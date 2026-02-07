package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Melbeck777/noq/infrastructure/repository"
	"github.com/Melbeck777/noq/internal/domain/entity"
	"github.com/Melbeck777/noq/internal/domain/valueobject"
)

func testConfigPath(t *testing.T) string {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}

	// テスト実行時のカレントがどこでも、リポジトリルートを起点にしたい場合:
	// その場合の repoRoot は 1つ上
	repoRoot := filepath.Dir(wd)

	dataDir := filepath.Join(repoRoot, "test", "data")
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}

	return filepath.Join(dataDir, "config.yml")
}

func removeIfExists(t *testing.T, path string) {
	t.Helper()
	_ = os.Remove(path) // 無ければ無視
}

func newValidConfig(t *testing.T) entity.Config {
	t.Helper()

	dbs, err := valueobject.NewNotionDatabases(map[string]string{
		"memo": "dummy-db-id",
	})
	if err != nil {
		t.Fatalf("NewNotionDatabases: %v", err)
	}
	alias, err := valueobject.NewDatabaseAlias("memo")
	if err != nil {
		t.Fatalf("NewDatabaseAlias: %v", err)
	}

	return entity.Config{
		MemoRoot:    "memo",
		ArticleRoot: "article",
		Notion: entity.NotionConfig{
			Token:         "dummy-token",
			DefaultMemoDB: *alias,
			Databases:     dbs,
		},
	}
}

// save->load
// expect : equal to save value
func TestConfigRepository_SaveLoad(t *testing.T) {
	path := testConfigPath(t)
	removeIfExists(t, path)

	repo := repository.NewConfigRepositoryImpl(path)

	want := newValidConfig(t)

	if err := repo.Save(want); err != nil {
		t.Fatalf("Save: %v", err)
	}

	got, err := repo.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if got.MemoRoot != want.MemoRoot {
		t.Fatalf("MemoRoot: got=%q want=%q", got.MemoRoot, want.MemoRoot)
	}
	if got.ArticleRoot != want.ArticleRoot {
		t.Fatalf("ArticleRoot: got=%q want=%q", got.ArticleRoot, want.ArticleRoot)
	}
	if got.Notion.Token != want.Notion.Token {
		t.Fatalf("Token: got=%q want=%q", got.Notion.Token, want.Notion.Token)
	}
	if string(got.Notion.DefaultMemoDB.Value()) != string(want.Notion.DefaultMemoDB.Value()) {
		t.Fatalf("DefaultMemoDB: got=%q want=%q", got.Notion.DefaultMemoDB, want.Notion.DefaultMemoDB)
	}

	// Databases の中身確認（valueobject 側に GetDBMap() がある前提）
	gotMap := got.Notion.Databases.GetDBMap()
	wantMap := want.Notion.Databases.GetDBMap()
	if len(gotMap) != len(wantMap) {
		t.Fatalf("Databases length: got=%d want=%d", len(gotMap), len(wantMap))
	}
	for k, v := range wantMap {
		gv, ok := gotMap[k]
		if !ok {
			t.Fatalf("Databases missing key: %q", k)
		}
		if gv != v {
			t.Fatalf("Databases[%q]: got=%q want=%q", k, gv, v)
		}
	}
}

// save->save
// expect : second save is falie
func TestConfigRepository_Save_AlreadyExists(t *testing.T) {
	path := testConfigPath(t)
	removeIfExists(t, path)

	repo := repository.NewConfigRepositoryImpl(path)
	cfg := newValidConfig(t)

	if err := repo.Save(cfg); err != nil {
		t.Fatalf("Save(1): %v", err)
	}

	// 2回目は失敗するはず
	if err := repo.Save(cfg); err == nil {
		t.Fatalf("Save(2) should fail when file exists")
	}
}

// save -> save overwrite
// expect : config.yml value is second save value.
func TestConfigRepository_SaveOverwrite(t *testing.T) {
	path := testConfigPath(t)
	removeIfExists(t, path)

	repo := repository.NewConfigRepositoryImpl(path)

	cfg1 := newValidConfig(t)
	cfg2 := newValidConfig(t)
	cfg2.MemoRoot = "memo2"
	cfg2.Notion.Token = "dummy-token-2"

	if err := repo.Save(cfg1); err != nil {
		t.Fatalf("Save(cfg1): %v", err)
	}

	if err := repo.SaveOverWrite(cfg2); err != nil {
		t.Fatalf("SaveOverwrite(cfg2): %v", err)
	}

	got, err := repo.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got.MemoRoot != cfg2.MemoRoot {
		t.Fatalf("MemoRoot after overwrite: got=%q want=%q", got.MemoRoot, cfg2.MemoRoot)
	}
	if got.Notion.Token != cfg2.Notion.Token {
		t.Fatalf("Token after overwrite: got=%q want=%q", got.Notion.Token, cfg2.Notion.Token)
	}
}
