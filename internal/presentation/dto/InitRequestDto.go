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
