package main

import "fmt"

func main() {
	// ### confusing slices
	x := make([]string, 0, 5)
	fmt.Println("x:", x)
	x = append(x, "a", "b", "c", "d")
	fmt.Println("x:", x)

	// full slice expression protects against confusing behavior due to append and memory sharing
	// s[<beginning (incl)> : <end (excl)> : <last position in the parent slice's capacity>]
	// sub-slice capacity = last position in the parent slice's capacity - beginning (incl)
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println("append to Y: i,j,k")
	y = append(y, "i", "j", "k")

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println("append to X: x")
	x = append(x, "x")

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println("append to Z: y")
	z = append(z, "y")

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("")
	fmt.Println("")

	// ### copying slices
	a := []int{1, 2, 3, 4}
	b := make([]int, 4)
	num := copy(b, a)
	fmt.Println(b, num)
}
