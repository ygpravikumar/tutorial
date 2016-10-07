package main

import "fmt"
import "math/cmplx"

func main() {
	var i int
	// the basic types are
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte alias for uint8
	// rune alias for int32 and represents a unicode code point
	// float32 float64
	// complex64 complex128

	// Variable declarations also can be factored like import statement
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Println("Hello there.", i)
	fmt.Println(ToBe)
	fmt.Println(MaxInt)
	fmt.Println(z)
}
