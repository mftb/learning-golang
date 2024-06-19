package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func rec_add(x, y int) int {
	if y > 0 {
		return 1 + rec_add(x, y-1)
	} else {
		return x
	}
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(rec_add(42, 13))
}
