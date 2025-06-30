package types

type FolderContentType int

const (
	FolderContentTypeFile   FolderContentType = 1
	FolderContentTypeFolder FolderContentType = 2
)

type FolderContent interface {
	GetName() string
	GetFolderContentType() FolderContentType
}
