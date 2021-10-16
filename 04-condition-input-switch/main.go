/*
Get the mark from user
use if else condition to rank the user
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your mark : ")
	scanner.Scan()

	mark, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Your mark is", mark)

	if mark < 35 {
		fmt.Println("You failed the Exam")
	} else if mark >= 35 && mark <= 50 {
		fmt.Println("You scored low in the Exam")
	} else if mark >= 35 && mark <= 50 {
		fmt.Println("You scored average in the Exam")
	} else if mark > 50 && mark <= 75 {
		fmt.Println("You scored well in the Exam")
	} else if mark > 75 && mark <= 95 {
		fmt.Println("You scored very well in the Exam")
	} else if mark > 95 && mark <= 100 {
		fmt.Println("You scored very high in the Exam")
	} else {
		fmt.Println("Enter a valid mark from 0 - 100")
	}

	switch {
	case mark < 35:
		fmt.Println("You failed the Exam")
	case mark >= 35 && mark <= 50:
		fmt.Println("You scored low in the Exam")
	case mark > 50 && mark <= 75:
		fmt.Println("You scored well in the Exam")
	case mark > 75 && mark <= 95:
		fmt.Println("You scored very well in the Exam")
	case mark > 95 && mark <= 100:
		fmt.Println("You scored very high in the Exam")
	default:
		fmt.Println("Enter a valid mark from 0 - 100")
	}

}
