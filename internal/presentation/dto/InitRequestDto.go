package dto

type InitRequestDto struct {
	MemoRoot      string
	ArticleRoot   string
	NotionToken   string
	DefaultMemoDB string
	Databases     map[string]string
}

type InitResponseDto struct {
	MemoRoot      string
	ArticleRoot   string
	NotionToken   string
	DefaultMemoDB string
	Databases     map[string]string
}

func NewInitRequestDto() (InitRequestDto, error) {
	m := make(map[string]string)
	return InitRequestDto{
		MemoRoot:      "",
		ArticleRoot:   "",
		NotionToken:   "",
		DefaultMemoDB: "",
		Databases:     m,
	}, nil
}
