// internal/domain/entity/NotionConfig.go
package entity

import (
	"github.com/Melbeck777/noq/internal/domain/valueobject"
)

type NotionConfig struct {
	Token         string
	DefaultMemoDB valueobject.DatabaseAlias
	Databases     valueobject.NotionDatabases
}

func NewNotionConfig(token, defaultMemoDb string, databases map[string]string) (NotionConfig, error) {
	dbAlias, err := valueobject.NewDatabaseAlias(defaultMemoDb)
	if err != nil {
		return NotionConfig{}, err
	}
	dbs, err := valueobject.NewNotionDatabases(databases)
	if err != nil {
		return NotionConfig{}, err
	}
	return NotionConfig{
		Token:         token,
		DefaultMemoDB: *dbAlias,
		Databases:     dbs,
	}, nil
}

func (n *NotionConfig) UpdateDefaultMemoDB(memo string) error {
	update_value, err := valueobject.NewDatabaseAlias(memo)
	if err != nil {
		return err
	}
	n.DefaultMemoDB = *update_value
	return nil
}
