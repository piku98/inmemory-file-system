package types

type Folder struct {
	Name     string
	Contents map[string]FolderContent
}

func (f *Folder) GetName() string {
	return f.Name
}

func (f *Folder) GetFolderContentType() FolderContentType {
	return FolderContentTypeFolder
}
