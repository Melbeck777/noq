package valueobject

type (
	DatabaseAlias string
	DatabaseID    string
)

type NotionDatabases struct {
	m map[DatabaseAlias]DatabaseID
}

func NewNotionDatabases(raw map[string]string) (NotionDatabases, error) {
	m := make(map[DatabaseAlias]DatabaseID, len(raw))
	for k, v := range raw {
		alias, err := NewDatabaseAlias(k)
		if err != nil {
			return NotionDatabases{}, err
		}

		id, err := NewDatabaseID(v)
		if err != nil {
			return NotionDatabases{}, err
		}

		m[alias] = id
	}
	return NotionDatabases{m: m}, nil
}

func NewDatabaseAlias(s string) (DatabaseAlias, error) {
	// 空文字禁止
	// 予約後の禁止
	return DatabaseAlias(s), nil
}

func NewDatabaseID(s string) (DatabaseID, error) {
	// 空文字禁止
	// uuid形式
	return DatabaseID(s), nil
}

// getの実装
func (d NotionDatabases) Get(alias DatabaseAlias) (DatabaseID, bool) {
	id, ok := d.m[alias]
	return id, ok
}
