package main

import (
	cmd "inmemory_file_system/command-line"
	filesystem "inmemory_file_system/file-system"
)

func main() {
	fileSystem := filesystem.NewInMemoryFileSystem()

	commandLine := cmd.NewCmd(fileSystem)
	commandLine.Run()

	// fileSystem.Root.Contents = []types.FolderContent{
	// 	&types.Folder{
	// 		Name: "users",
	// 		Contents: []types.FolderContent{
	// 			&types.File{
	// 				Name:    "some file",
	// 				Type:    "txt",
	// 				Size:    0,
	// 				Content: []byte{},
	// 			},
	// 			&types.Folder{
	// 				Name: "sourav",
	// 				Contents: []types.FolderContent{
	// 					&types.File{
	// 						Name:    "sourav file",
	// 						Type:    "txt",
	// 						Size:    0,
	// 						Content: []byte{},
	// 					},
	// 					&types.Folder{
	// 						Name: "desktop",
	// 						Contents: []types.FolderContent{
	// 							&types.File{
	// 								Name:    "some pdf",
	// 								Type:    "txt",
	// 								Size:    0,
	// 								Content: []byte{},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// fmt.Println(fileSystem.GetCurrentFolder())
	// err := fileSystem.Mkdir("~/users/pp/balhhh/pookie")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = fileSystem.Cd("~/users/pp/balhhh/pookie")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = fileSystem.Cd("~/users")
	// err = fileSystem.Cd("sourav/desktop")
	// contents := fileSystem.Ls()
	// fmt.Println(fileSystem.GetCurrentFolder(), contents)
	// content, err := fileSystem.Open("~/users/sourav/desktop/some pdf")
	// if err != nil {
	// 	fmt.Println("OPEN ERROR", err)
	// }
	// fmt.Println(content)
}
