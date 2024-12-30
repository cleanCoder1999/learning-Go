package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// ### - exercise 1: check what escapes to the heap => build with: go build -gcflags="-m"
	// -gcflags="-m" ... prints out which values escape to the heap
	{
		p := MakePerson("Ken", "Thompson", 88)
		fmt.Println(p) // why does p escape to the heap here??

		pp := MakePersonPointer("Dennis", "Ritchie", 0)
		fmt.Println(pp)
		fmt.Println("")
	}

	// ### - exercise 2: slice: modify content and grow
	{
		sl := []string{"Ken", "Thompson", "Dennis", "Ritchie"}
		fmt.Println("before update:", sl)
		UpdateSlice(sl, "Rob")
		fmt.Println("after update:", sl)
		fmt.Println("")

		fmt.Println("before grow:", sl)
		GrowSlice(sl, "Pike")
		fmt.Println("after grow:", sl)
		fmt.Println("")
	}

	// ### - exercise 3: tune the garbage collector
	// GODEBUG=gctrace=1 ... allows you to see when garbage collection happens
	// set it in the terminal with 'export GODEBUG=gctrace=1'
	// GOGC ... allows to manipulate the rate of how often a garbage collection cycle happens
	// GOGC=100 is the default
	// => >100 reduces the rate
	// => <100 increases the rate
	{
		p := Person{FirstName: "Dennis", LastName: "Ritchie", Age: 0}
		people := make([]Person, 0)

		start := time.Now()
		fmt.Println(start)
		for i := 0; i < 10_000_000; i++ {
			people = append(people, p)
		}
		end := time.Now()
		fmt.Println(end)
		fmt.Println("duration:", end.Sub(start))
	}

}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{FirstName: firstName, LastName: lastName, Age: age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age} // why does this escape to the heap??
}

func UpdateSlice(sl []string, s string) {
	last := len(sl) - 1
	sl[last] = s
	fmt.Println(sl)
}

func GrowSlice(sl []string, s string) {
	sl = append(sl, s)
	fmt.Println(sl)
}
