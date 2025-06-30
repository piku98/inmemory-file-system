package cmd

import (
	"bufio"
	"fmt"
	filesystem "inmemory_file_system/file-system"
	"os"
	"strings"
)

type Cmd struct {
	FS filesystem.FileSystem
}

func NewCmd(fileSystem filesystem.FileSystem) *Cmd {
	return &Cmd{
		FS: fileSystem,
	}
}

func (t *Cmd) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to InMemoryFileSystem Terminal!")
	for {
		fmt.Printf("[%s]$ ", t.FS.GetCurrentFolder())
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		switch cmd {
		case "exit":
			return
		case "ls":
			contents := t.FS.Ls()
			fmt.Println(strings.Join(contents, " "))
		case "cd":
			if len(args) < 2 {
				fmt.Println("Usage: cd <path>")
				continue
			}
			if err := t.FS.Cd(args[1]); err != nil {
				fmt.Println("cd error:", err)
			}
		case "mkdir":
			if len(args) < 2 {
				fmt.Println("Usage: mkdir <path>")
				continue
			}
			if err := t.FS.Mkdir(args[1]); err != nil {
				fmt.Println("mkdir error:", err)
			}
		case "open":
			if len(args) < 2 {
				fmt.Println("Usage: open <path>")
				continue
			}
			content, err := t.FS.Open(args[1])
			if err != nil {
				fmt.Println("open error:", err)
			} else {
				fmt.Println(string(content))
			}
		case "showtree":
			if len(args) > 1 {
				fmt.Println("no args in command")
				continue
			}
			fmt.Println(t.FS.ShowTree())
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}
