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

func NewNotionConfig(token, defaultMemoDb string, databases map[string]string) *NotionConfig {
	dbAlias, _ := valueobject.NewDatabaseAlias(defaultMemoDb)
	dbs, _ := valueobject.NewNotionDatabases(databases)
	return &NotionConfig{
		Token:         token,
		DefaultMemoDB: *dbAlias,
		Databases:     dbs,
	}
}

func (n *NotionConfig) UpdateDefaultMemoDB(memo string) error {
	update_value, err := valueobject.NewDatabaseAlias(memo)
	if err != nil {
		return err
	}
	n.DefaultMemoDB = *update_value
	return nil
}
