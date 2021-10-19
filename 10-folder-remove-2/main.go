package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error - ", e)
		os.Exit(1)
	}
}

func main() {

	//err := os.RemoveAll("tmp")
	//scheck(err)

	dir, err := ioutil.ReadDir("tmp")
	check(err)
	for _, d := range dir {
		//fmt.Println(d.Name())
		fmt.Println(path.Join("tmp", d.Name()))
		os.RemoveAll(path.Join("tmp", d.Name()))
		//fmt.Println(path.Join([]string{"tmp", d.Name()}...))
		//os.RemoveAll(path.Join([]string{"tmp", d.Name()}...))
	}

}
