package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error - ", e)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Files with Golang")

	// create a file
	file, err := os.Create("akilan.txt")

	check(err)
	defer file.Close() // it will execute at last

	// write a data to a file

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter data > ")
		scanner.Scan()

		data := scanner.Text()

		if data == "quit" {
			break
		}

		file.WriteString(data + "\n")

	}

	// read a file
	byte_data, err := os.ReadFile(file.Name())

	check(err)

	fmt.Println(string(byte_data))

	// info about file

	file_info, err := os.Stat(file.Name())
	check(err)
	fmt.Printf("Is Directory ? %t \n", file_info.IsDir())
	fmt.Printf("Size of the file is %d \n", file_info.Size())

	// Renaming a file
	err = os.Rename(file.Name(), "family.txt")
	check(err)

	// removing a file
	//err = os.Remove(file.Name())
	//check(err)

}
