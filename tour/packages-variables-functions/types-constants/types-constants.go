package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

func main() {
	var x, y int = 3, 4
	bigint := 10_000
	bigfloat := 10_000.1
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// for some reason this yields something that is not the string representation of 10_000
	var s string = string(bigint)
	fmt.Println(x, y, z, s, bigint)
	fmt.Printf("bigint is of type %T\n", bigint)
	fmt.Printf("bigfloat is of type %T\n", bigfloat)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
