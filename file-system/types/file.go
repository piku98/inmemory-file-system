package types

type File struct {
	Name    string
	Size    int64
	Type    string
	Content []byte
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) GetFolderContentType() FolderContentType {
	return FolderContentTypeFile
}
