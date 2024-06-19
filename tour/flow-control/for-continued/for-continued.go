package main

import "fmt"

func main() {
	sum := 1
	// auto formatter removed the empty init and post statements
	// for is golang's while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
