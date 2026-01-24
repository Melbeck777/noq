package entity

type Config struct {
	// memo_root in config.yml
	MemoRoot    string
	ArticleRoot string
	Notion      NotionConfig
}
