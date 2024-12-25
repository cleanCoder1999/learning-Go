package main

import "fmt"

func main() {
	// ### confusing slices
	x := make([]string, 0, 5)
	fmt.Println("x:", x)
	x = append(x, "a", "b", "c", "d")
	fmt.Println("x:", x)

	y := x[:2]
	z := x[2:]
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
}
