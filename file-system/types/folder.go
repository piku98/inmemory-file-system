package types

type Folder struct {
	Name     string
	Contents []FolderContent
}

func (f *Folder) GetName() string {
	return f.Name
}

func (f *Folder) GetFolderContentType() FolderContentType {
	return FolderContentTypeFolder
}
