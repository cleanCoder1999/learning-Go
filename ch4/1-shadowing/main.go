package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(true)

	// Go contains only 25 keywords;
	// predeclared identifiers such as int, string, bool, true or false are not part of those 25 keywords
	// those are part of Go's "universal block"
	// the universal block is the block that contains all other blocks of a Go program
	true := "shadowing bool's 'true'"
	fmt.Println(true)

	var b bool
	fmt.Println("is of type:", reflect.TypeOf(b))

	bool := "now, bool is a string lol"
	fmt.Println(bool)
}
