package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func cleanupFiles(p string, file_info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if file_info.IsDir() {
		//fmt.Println("Skipping as it a folder")
	} else {
		fmt.Println("Removing file ", p)
		err = os.Remove(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func cleanupFolders(p string, file_info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	// removing empty folder
	if file_info.IsDir() && p != "tmp" {
		fmt.Println("Removing folder ", p)

		err = os.Remove(p)
		if err != nil {
			return err
		}

	}

	return nil
}

func main() {

	fmt.Println("Removes all the files and folders")

	filepath.Walk("tmp", cleanupFiles)
	err := filepath.Walk("tmp", cleanupFolders)

	if err != nil {
		fmt.Println(err)
	}
}
