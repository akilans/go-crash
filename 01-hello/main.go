package main

import (
	"fmt"

	"github.com/akilans/go-crash/hello-module"
)

func double(num int) int {
	return num + num
}

func main() {
	hello.Greet()
	fmt.Println(double(2))
}

// go test -v - shows success and failure report
// go test -v -cover - shows coverage report as well
// go test -v -cover -coverprofile=c.out
// go tool cover -html=c.out -o coverage.html
