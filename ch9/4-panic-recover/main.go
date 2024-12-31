package main

import "fmt"

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}

func div60(i int) {

	// NOTE: recover() must be called from within a defer because once a panic happens, only deferred functions are run
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	fmt.Println(60 / i)
}