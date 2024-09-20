package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func setTemplatePath(args ...string) string {
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir)
	for _, arg := range args {
		filePath += string(filepath.Separator) + arg
	}

	return filePath
}

func main() {
	fmt.Println(setTemplatePath("first", "second", "third", "fourth"))
}
