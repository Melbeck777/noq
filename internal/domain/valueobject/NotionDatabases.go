// internal/domain/valuobject/NotionDatabases.go
package valueobject

import (
	"fmt"
)

type DatabaseAlias struct {
	ValueObject[string]
}

func NewDatabaseAlias(value string) (*DatabaseAlias, error) {
	if value == "" {
		return nil, fmt.Errorf("DatabaseAlias cannot be empty")
	}
	return &DatabaseAlias{NewValueObject[string](value)}, nil
}

type DatabaseId struct {
	ValueObject[string]
}

func NewDatabaseId(value string) (*DatabaseId, error) {
	if value == "" {
		return nil, fmt.Errorf("DatabaseId cannot be empty")
	}
	return &DatabaseId{NewValueObject[string](value)}, nil
}

type NotionDatabases struct {
	m map[string]string
}

func (d NotionDatabases) GetDBMap() map[string]string {
	return d.m
}

func NewNotionDatabases(raw map[string]string) (NotionDatabases, error) {
	m := make(map[string]string, len(raw))
	for k, v := range raw {
		if _, err := NewDatabaseAlias(k); err != nil {
			return NotionDatabases{}, err
		}

		if _, err := NewDatabaseId(v); err != nil {
			return NotionDatabases{}, err
		}
		m[k] = v
	}
	return NotionDatabases{m: m}, nil
}

// getの実装
func (d NotionDatabases) Get(alias string) (string, bool) {
	id, ok := d.m[alias]
	return id, ok
}
