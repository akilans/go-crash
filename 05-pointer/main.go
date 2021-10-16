package main

import "fmt"

func incByValue(age int) {
	age++
}

func incByRef(age *int) {
	fmt.Println("It is a memory address", age)
	*age++
}

func main() {

	age := 34

	//pass by value
	incByValue(age)
	fmt.Println("Your age after incByValue", age)

	//pass by ref
	incByRef(&age)
	fmt.Println("Your age after incByRef", age)

	numbers1 := []int{1, 2, 3, 4}

	numbers2 := numbers1 // Point to the same address

	numbers2[1] = 100

	fmt.Println(numbers1) //[1 100 3 4]

}
