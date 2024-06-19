package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// really like how you can extract different data in the printf with the %x inserts
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// interesting group var notation
	var (
		ToBeLocal   bool       = false
		MaxIntLocal uint64     = 1<<64 - 1
		zLocal      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBeLocal, ToBeLocal)
	fmt.Printf("Type: %T Value: %v\n", MaxIntLocal, MaxIntLocal)
	fmt.Printf("Type: %T Value: %v\n", zLocal, zLocal)
}
