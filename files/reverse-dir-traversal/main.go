package main

import (
	"fmt"
	"os"
)

func listFilesAndDirs(path string, indent string) {
	dir, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening directory %s: %v\n", path, err)
		return
	}
	defer dir.Close()

	entries, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", path, err)
		return
	}

	// Iterate over the entries
	for _, entry := range entries {
		// Print the entry name and type
		fmt.Printf("%s%s\n", indent, entry.Name())

		if entry.IsDir() {
			subDirPath := path + "/" + entry.Name()   // Manually construct the subdirectory path
			listFilesAndDirs(subDirPath, indent+"  ") // Add indentation for subdirectories
		}
	}
}

func main() {
	root := "../"
	fmt.Println("Files and directories in the folder:")
	listFilesAndDirs(root, "")
}
