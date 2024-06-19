package main

import "fmt"

// tuple like in python
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// tuple unpacking like in python
	a, b := swap("hello", "world")
	// dissappointed that x, y := y, x does not work
	fmt.Println(a, b)
}
