// internal/domain/entity/NotionConfig.go
package entity

import (
	"github.com/Melbeck777/noq/internal/domain/valueobject"
)

type NotionConfig struct {
	Token         string
	DefaultMemoDB *valueobject.DatabaseAlias
	Databases     valueobject.NotionDatabases
}

func NewNotionConfig(token, defaultMemoDb string, databases map[string]string) (NotionConfig, error) {
	var dbAlias *valueobject.DatabaseAlias
	if defaultMemoDb != "" {
		tmpAlias, err := valueobject.NewDatabaseAlias(defaultMemoDb)
		if err != nil {
			return NotionConfig{}, err
		}
		dbAlias = tmpAlias
	}
	dbs, err := valueobject.NewNotionDatabases(databases)
	if err != nil {
		return NotionConfig{}, err
	}
	return NotionConfig{
		Token:         token,
		DefaultMemoDB: dbAlias,
		Databases:     dbs,
	}, nil
}

func (n *NotionConfig) UpdateDefaultMemoDB(memo string) error {
	updateAlias, err := valueobject.NewDatabaseAlias(memo)
	if err != nil {
		return err
	}
	n.DefaultMemoDB = updateAlias
	return nil
}
