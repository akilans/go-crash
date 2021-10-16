/*
Function returning another function
Function can be assigned to a variable
*/
package main

import "fmt"

//  function is returning another function
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// function can be assigned into a variable
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
