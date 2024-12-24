package main

import "fmt"

func main() {
	// ### explicit type conversion
	i := 20
	var f = float64(i)

	fmt.Println(i)
	fmt.Println(f)

	// ### untyped constant
	const value = 10

	i = value
	f = value
	fmt.Println(i)
	fmt.Println(f)

	// ### max values and overflow
	var b byte = 127
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807
	fmt.Printf("max: %d, max+1: %d\n", b, b+1)
	fmt.Printf("max: %d, max+1: %d\n", smallI, smallI+1)
	fmt.Printf("max: %d, max+1: %d\n", bigI, bigI+1)

}
