package main

import (
	cmd "inmemory_file_system/command-line"
	filesystem "inmemory_file_system/file-system"
)

func main() {
	fileSystem := filesystem.NewInMemoryFileSystem()

	commandLine := cmd.NewCmd(fileSystem)
	commandLine.Run()
}
