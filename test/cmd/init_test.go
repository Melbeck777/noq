package cmd

import (
	"os"
	"path/filepath"
	"testing"
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
