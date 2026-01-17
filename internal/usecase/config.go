package usecase

type Config interface {
	MemoRootPath() string
	ArticleRootPath() string
	NotionMemoDBId() string
	NotionMemoDBListId() map[string]string
}
