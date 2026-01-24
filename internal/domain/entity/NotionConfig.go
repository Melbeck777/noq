package entity

import (
	"github.com/Melbeck777/noq/internal/domain/valueobject"
)

type NotionConfig struct {
	Token         string
	DefaultMemoDB valueobject.DatabaseAlias
	Databases     valueobject.NotionDatabases
}
