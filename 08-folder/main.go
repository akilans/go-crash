package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error - ", e)
		os.Exit(1)
	}
}

func listFileFolder(p string, info os.FileInfo, err error) error {
	check(err)
	fmt.Println(" ", p, info.IsDir())
	return nil
}

func main() {
	fmt.Println("Folders with Golang")

	err := os.MkdirAll("akilan/pooma/inba/iniya/", 0777)
	check(err)

	// remove only empty dir
	//err = os.Remove("akilan/pooma/inba/iniya/")
	//check(err)

	// returns only childs,  wont show childs child
	dirs, err := os.ReadDir("akilan")

	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}
	//Pooma

	filepath.Walk("akilan", listFileFolder)
	/*
	  akilan true
	  akilan/pooma true
	  akilan/pooma/inba true
	  akilan/pooma/inba/hello.text false
	*/

}
