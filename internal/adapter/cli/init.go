package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type InitInput struct {
	MemoRoot    string
	ArticleRoot string
}

func CollectInitInput(memoRootFlag, articleRootFlag string) (InitInput, error) {
	defMemo, defArticle, err := defaultRoots()
	if err != nil {
		return InitInput{}, err
	}

	memoRoot := memoRootFlag
	articleRoot := articleRootFlag

	in := bufio.NewReader(os.Stdin)

	if strings.TrimSpace(memoRoot) == "" {
		memoRoot, err = askPathWithDefault(in, "memo_root", defMemo)
		if err != nil {
			return InitInput{}, err
		}
	}
	if strings.TrimSpace(articleRoot) == "" {
		articleRoot, err = askPathWithDefault(in, "article_root", defArticle)
		if err != nil {
			return InitInput{}, err
		}
	}

	memoRoot, err = normalizePath(memoRoot)
	if err != nil {
		return InitInput{}, err
	}
	if err := ensureDir(memoRoot); err != nil {
		return InitInput{}, err
	}
	articleRoot, err = normalizePath(articleRoot)
	if err != nil {
		return InitInput{}, err
	}
	if err := ensureDir(articleRoot); err != nil {
		return InitInput{}, err
	}

	if memoRoot == "" {
		return InitInput{}, fmt.Errorf("memo_root is required")
	}
	if articleRoot == "" {
		return InitInput{}, fmt.Errorf("article_root is required")
	}

	return InitInput{
		MemoRoot:    memoRoot,
		ArticleRoot: articleRoot,
	}, nil
}

func ensureDir(path string) error {
	if err := os.MkdirAll(path, 0o755); err != nil {
		return fmt.Errorf("failed to create directory: %s: %w", path, err)
	}
	return nil
}

func defaultRoots() (string, string, error) {
	memo, err := defaultMemoRoot()
	if err != nil {
		return "", "", err
	}
	article, err := defaultArticleRoot()
	if err != nil {
		return "", "", err
	}
	return memo, article, nil
}

func askPathWithDefault(in *bufio.Reader, label, def string) (string, error) {
	fmt.Printf("%s (default: %s):", label, def)
	line, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}

	v := strings.TrimSpace(line)
	if v == "" {
		return def, nil
	}

	return v, nil
}

func normalizePath(p string) (string, error) {
	if p == "" {
		return "", nil
	}

	if p == "~" || strings.HasPrefix(p, "~/") || strings.HasPrefix(p, `~\`) {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		rest := strings.TrimPrefix(p, "~")
		rest = strings.TrimPrefix(rest, "/")
		rest = strings.TrimPrefix(rest, `\`)
		p = filepath.Join(home, rest)
	}

	if !filepath.IsAbs(p) {
		abs, err := filepath.Abs(p)
		if err != nil {
			return "", err
		}
		p = abs
	}
	return filepath.Clean(p), nil
}

func defaultMemoRoot() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "noq", "memo"), nil
}

func defaultArticleRoot() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "noq", "article"), nil
}
