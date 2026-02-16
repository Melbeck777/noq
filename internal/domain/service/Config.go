package service

import (
	"fmt"

	"github.com/Melbeck777/noq/internal/domain/entity"
)

func EmptyCheck(target, title string, errorType error) error {
	if target == "" {
		return fmt.Errorf("%w: %s is empty\n", errorType, title)
	}
	return nil
}

type DiffValue struct {
	Key      string
	olValue  string
	newValue string
}

type DuplicateKeys struct {
	firstKey     string
	duplicateKey string
}

type MergeResult struct {
	MergeConfig   entity.Config
	DiffValues    []DiffValue
	DuplicateKeys []DuplicateKeys
}

func Merge(old_entity, new_entity entity.Config) (MergeResult, error) {
	var res entity.Config
	var diff_values []DiffValue
	var duplicate_keys []DuplicateKeys
	if old_entity.MemoRoot != new_entity.MemoRoot {
		diff_values = append(diff_values, DiffValue{
			"MemoRoot",
			old_entity.MemoRoot,
			new_entity.MemoRoot,
		})
	} else {
		res.MemoRoot = new_entity.MemoRoot
	}

	if old_entity.ArticleRoot != new_entity.ArticleRoot {
		diff_values = append(diff_values, DiffValue{
			"ArticleRoot",
			old_entity.ArticleRoot,
			new_entity.ArticleRoot,
		})
	} else {
		res.ArticleRoot = new_entity.ArticleRoot
	}

	if old_entity.Notion.Token != new_entity.Notion.Token {
		diff_values = append(diff_values, DiffValue{
			"Notion.Token",
			old_entity.Notion.Token,
			new_entity.Notion.Token,
		})
	} else {
		res.Notion.Token = new_entity.Notion.Token
	}

	if old_entity.Notion.DefaultMemoDB != new_entity.Notion.DefaultMemoDB {
		diff_values = append(diff_values, DiffValue{
			"Notion.DefaultMemoDB",
			old_entity.Notion.DefaultMemoDB.Value(),
			new_entity.Notion.DefaultMemoDB.Value(),
		})
	} else {
		res.Notion.DefaultMemoDB = new_entity.Notion.DefaultMemoDB
	}
	// 同じキーが入っているかを調べる > 入ってたら排除かどうか聞くため
	// 同じキーに対して異なる値が入っているかを調べる > diff出すため
	var value_keys map[string]string
	for k, v := range new_entity.Notion.Databases.GetDBMap() {
		id, ok := old_entity.Notion.Databases.Get(k)
		// 既存比較
		if ok {
			if v != id {
				diff_values = append(diff_values, DiffValue{
					"Notion.Dabases.%s" + k,
					id,
					v,
				})
			}
		}

		// 重複検証
		if old_k, ok := value_keys[v]; ok {
			duplicate_keys = append(duplicate_keys, DuplicateKeys{
				old_k,
				k,
			})
		} else {
			value_keys[v] = k
		}
	}

	return MergeResult{res, diff_values, duplicate_keys}, nil
}
