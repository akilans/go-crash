package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	Name     string
	Age      int
	Location string
}

// Normal function
func sayHello(name string) (string, error) {
	if name != "Akilan" {
		return "", errors.New("Error - You are not Akilan")
	}
	return "Hello " + name, nil
}

// Receiver function - It can be called only by Employee struct
func (e Employee) greetEmployee() string {
	greet := fmt.Sprintf("Hello %s, You are %d years old and live in %s", e.Name, e.Age, e.Location)
	return greet
}

// variadic function

func sum(nums ...int) int {
	result := 0

	for _, num := range nums {
		result = result + num
	}

	return result

}

func main() {
	fmt.Println("Welcome to functions")
	greet, err := sayHello("Akila")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greet)
	}

	akilan := Employee{"Akilan", 34, "Bangalore"}
	fmt.Println(akilan.greetEmployee())

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
}
