package main

import "fmt"

func main() {
	fmt.Println("Variadic Input Params:")
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
}

func addTo(base int, values ...int) []int {
	out := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		out[i] = base + values[i]
	}

	return out
}
