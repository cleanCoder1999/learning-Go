package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// ### - exercise 1 and 2
	{
		epoch := time.Now().Unix()
		r := rand.New(rand.NewSource(epoch))

		rands := make([]int, 100)
		for i := range rands {
			rands[i] = r.Int()
		}

		for _, v := range rands {
			if v%2 == 0 {
				fmt.Println("Two")
				continue
			}
			if v%3 == 0 {
				fmt.Println("Three")
				continue
			}
			if v%2 == 0 && v%3 == 0 {
				fmt.Println("Six")
				continue
			}

			fmt.Println("Never mind.")
		}
		fmt.Println("")
	}

	// ### - exercise 3 (shadowing)
	{
		var total int
		for i := 0; i < 10; i++ {
			total = total + i
			fmt.Println(total)
		}
		fmt.Println(total)
	}
}
