package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	file, err := os.Create("text.txt")
	handleError(err)
	defer file.Close()

	_, err = file.WriteString("This is a test file.\n")
	handleError(err)

	file, err = os.Open("main.go")
	handleError(err)

	stat, err := file.Stat()
	handleError(err)
	fmt.Println("File opened:", stat.Size())

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	handleError(err)

	dir, err := os.Open("../../chapter-7")
	handleError(err)
	defer dir.Close()

	fileInfos, err := dir.ReadDir(-1)
	handleError(err)

	for _, fi := range fileInfos {
		if fi.IsDir() {
			fmt.Printf("the %s is a Folder\n", fi.Name())
		} else {
			fmt.Printf("the %s isn't Folder\n\n", fi.Name())
		}
	}

	filepath.Walk("../../chapter-7", func(path string, info os.FileInfo, err error) error {
		handleError(err)
		fmt.Println(path)
		if !info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	})
}
