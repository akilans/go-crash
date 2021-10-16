/*
var int, string, slice, struct
for loop
while loop
*/
package main

import "fmt"

// standard way to define var
var fname string

//fname = "Akilan" - not possible to define outside func
var age int = 33

// const value
const native string = "tenkasi"

// Create own data structure
type Friend struct {
	name     string
	age      int
	location string
}

// Map example

var likes = map[string]string{
	"sports":   "kabaddi",
	"location": "Bangalore",
	"love":     "Nature",
}

func main() {
	fname = "Akilan"
	location := "Bangalore"
	fmt.Printf("Your native is %s\n", native)
	// native = "Bangalore" // cannot reassign const
	fmt.Printf("Hello %s your age is %d and you are in %s\n", fname, age, location)
	//Hello Akilan your age is 33 and you are in Bangalore

	var friends []Friend // slice(no need to specify size)

	var alex = Friend{"Alex", 33, "Chennai"}
	var jegan = Friend{"Jegan", 32, "Coimbatore"}

	friends = append(friends, alex)
	friends = append(friends, jegan)

	fmt.Println(friends)
	//[{Alex 33 Chennai} {Jegan 32 Coimbatore}]

	for index, friend := range friends {
		fmt.Printf("%d - Name %s Age %d Location %s \n", index+1, friend.name, friend.age, friend.location)
	}
	/*
		1 - Name Alex Age 33 Location Chennai
		2 - Name Jegan Age 32 Location Coimbatore
	*/

	for key, value := range likes {
		fmt.Printf("%s =====> %s \n", key, value)
	}
	/*
		love =====> Nature
		sports =====> kabaddi
		location =====> Bangalore
	*/
	i := 0
	// while loop
	for i < 5 {
		fmt.Println(i)
		i++
	}

}
