package main

import "fmt"

func main() {
	// defers the execution of a function until the surrounding function returns
	// the deferred call's arguments are evaluated immediately
	// but the function call is not executed until the surrounding function returns
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	// deferred function calls are pushed onto a stack
	// when a function returns, its deferred calls are executed in last-in-first-out order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
