package filesystem

import (
	"encoding/json"
	"fmt"
	"inmemory_file_system/file-system/types"
	"strings"
)

type FileSystem interface {
	GetCurrentFolder() string
	Cd(path string) error
	Ls() []string
	Mkdir(path string) error
	Open(path string) ([]byte, error)
	ShowTree() string
}

type InMemoryFileSystem struct {
	Root          *types.Folder
	currentFolder *types.Folder
}

func NewInMemoryFileSystem() *InMemoryFileSystem {
	rootFolder := &types.Folder{
		Name:     "~",
		Contents: []types.FolderContent{},
	}
	return &InMemoryFileSystem{
		Root:          rootFolder,
		currentFolder: rootFolder,
	}
}

func (fs *InMemoryFileSystem) GetCurrentFolder() string {
	return fs.currentFolder.Name
}

func (fs *InMemoryFileSystem) traverse(folderStrings []string, createNodeIfNotPresent bool) (*types.Folder, error) {

	folderStringsIdx := 0
	currentFolder := fs.currentFolder

	if folderStrings[0] == "~" {
		currentFolder = fs.Root
		folderStringsIdx += 1
	}

	for folderStringsIdx < len(folderStrings) {

		folderFound := false
		for _, content := range currentFolder.Contents {
			if content.GetFolderContentType() == types.FolderContentTypeFile ||
				content.GetName() != folderStrings[folderStringsIdx] {
				continue
			}
			currentFolder = content.(*types.Folder)
			folderFound = true
		}

		if !folderFound {
			if createNodeIfNotPresent {
				folder := &types.Folder{
					Name:     folderStrings[folderStringsIdx],
					Contents: []types.FolderContent{},
				}
				currentFolder.Contents = append(currentFolder.Contents, folder)
				currentFolder = folder
			} else {
				return nil, fmt.Errorf("folder not found")
			}
		}
		folderStringsIdx += 1
	}
	return currentFolder, nil
}

func (fs *InMemoryFileSystem) Cd(path string) error {
	folderStrings := strings.Split(path, "/")
	currentFolder, err := fs.traverse(folderStrings, false)
	if err != nil {
		return err
	}
	fs.currentFolder = currentFolder
	return nil
}

func (fs *InMemoryFileSystem) Ls() []string {
	contentStrings := make([]string, len(fs.currentFolder.Contents))

	for idx, content := range fs.currentFolder.Contents {
		contentStrings[idx] = content.GetName()
	}

	return contentStrings
}

func (fs *InMemoryFileSystem) Mkdir(path string) error {
	folderStrings := strings.Split(path, "/")
	_, err := fs.traverse(folderStrings, true)
	if err != nil {
		return err
	}
	return nil
}

func (fs *InMemoryFileSystem) Open(path string) ([]byte, error) {
	folderStrings := strings.Split(path, "/")
	currentFolder := fs.currentFolder
	fileName := path
	if len(folderStrings) > 1 {
		folderPath := folderStrings[:len(folderStrings)-1]
		folder, err := fs.traverse(folderPath, false)
		if err != nil {
			return nil, err
		}
		fileName = folderStrings[len(folderStrings)-1]
		currentFolder = folder
	}
	var file *types.File
	for _, content := range currentFolder.Contents {
		if content.GetFolderContentType() != types.FolderContentTypeFile || content.GetName() != fileName {
			continue
		}
		file = content.(*types.File)
	}
	if file != nil {
		return file.Content, nil
	}
	return nil, fmt.Errorf("could not locate file")
}

func (fs *InMemoryFileSystem) ShowTree() string {
	tree, _ := json.MarshalIndent(fs.Root, " ", " ")
	return string(tree)
}
