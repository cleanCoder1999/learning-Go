package main

import "fmt"

func main() {

	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evens := Filter(ints, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(evens)

	doubled := Map(ints, func(i int) int {
		return i * 2
	})
	fmt.Println(doubled)

	sum := Reduce(ints, 0, func(a, b int) int {
		return a + b
	})
	fmt.Println(sum)
}

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))

	for i, v := range s {
		r[i] = f(v)
	}

	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer

	for _, v := range s {
		r = f(r, v)
	}

	return r
}
