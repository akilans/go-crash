package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Testing concurrency")

	order_res := make(chan string)
	pay_res := make(chan string)

	go process("Ordering...", order_res)
	go process("Paying...", pay_res)

	// 10 times it receives
	for i := 0; i < 10; i++ {
		select {
		case o_res := <-order_res:
			fmt.Println(o_res)
		case p_res := <-pay_res:
			fmt.Println(p_res)
		}
	}

	/*
		for o_res := range order_res {
			fmt.Println(o_res)
		}

		for p_res := range pay_res {
			fmt.Println(p_res)
		}
	*/

}

func process(action string, c chan string) {
	defer close(c) // It will call last and best way to close channel
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		c <- action
	}
}
