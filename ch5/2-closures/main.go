package main

import (
	"fmt"
	"sort"
)

func main() {

	// ### - closure as parameter
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Ken", "Thompson", 88},
		{"Dennis", "Ritchie", 0},
		{"Rob", "Pike", 60},
	}

	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
	fmt.Println("")

	// ### - closure as return value
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
