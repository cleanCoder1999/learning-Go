package main

import "fmt"

func main() {
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}

	// method of type Inner
	// NOTE: 	even if Outer contains a method IntPrinter like Inner,
	//			the method on the embedded field is invoked,
	//			not the method on the containing struct
	fmt.Println(o.Double())
}

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner // embedded field
	S     string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}
