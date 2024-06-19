package main

import "fmt"

var c, python, java bool

// initialize variables
var j, k int = 1, 2

func main() {
	var i int
	// type is defined during initialization
	var c_init, python_init, java_init = true, false, "no!"
	my_var := 3
	// 0 false false false
	// differently than C, looks like golang has default init values for variables
	fmt.Println(i, c, python, java)
	// displays the initialized values
	fmt.Println(j, k, c_init, python_init, java_init)
	fmt.Println(my_var)
}
