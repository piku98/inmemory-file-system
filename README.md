## In memory file system with using trie

This is an attempt to build an in-memory file system which supports terminal commands and a few more like `cd`, `mkdir`, `ls`, `open`, `showtree`.
I tried to build the folder structure using trie datastructure.

## Usage

1. `cd <path>`: change directory
2. `mkdir <path>`: recursively create directories
3. `ls`: show contents in a folder
4. `open <path>`: returns the contents in a file as bytes
5. `showtree`: shows the entire tree structure of the filesystem.

## Example

<img width="393" alt="Screenshot 2025-06-30 at 5 24 58 PM" src="https://github.com/user-attachments/assets/bced2d7f-f21e-4dc3-8c16-38667b9297c0" />
