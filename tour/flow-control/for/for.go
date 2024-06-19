package main

import "fmt"

func main() {
	fn_0 := 0
	fn_1 := 1
	for i := 1; i <= 8; i++ {
		// := declares a new variable
		tmp := fn_1
		// = assigns value
		fn_1 = fn_1 + fn_0
		fn_0 = tmp
	}
	fmt.Println(fn_1)
}
