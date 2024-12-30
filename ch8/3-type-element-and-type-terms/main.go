package main

import (
	"errors"
	"fmt"
)

func main() {

	// ### - check if it works fine with values that are out of int range
	{
		var a uint = 18_345_678_999_567_534_563
		var b uint = 9_345_111_999_567_456_563
		fmt.Println(divAndRemainderGeneric(a, b))
	}

	// ### - provoke error due to exact match of type terms (default)
	{
		type MyInt int
		var myA MyInt = 10
		var myB MyInt = 20
		//fmt.Println(divAndRemainderGeneric(myA, myB)) // returns: MyInt does not satisfy Integer (possibly missing ~ for int in Integer)
		fmt.Println(myA, myB)
	}

	// ### - avoid error by allowing non-exact match of type terms
	{
		type MyInt int
		var myA MyInt = 10
		var myB MyInt = 20
		fmt.Println(divAndRemainderGenericAllowsNonExactMatch(myA, myB)) // works fine
		fmt.Println(myA, myB)
	}
}

// EXPLANATION FOR THE CODE BELOW:
// the function divAndRemainder() works fine for params of type int
//
// However, if you want it work for any integer type by making it a generic function,
// it is necessary to use a type element (which is composed of one or more type terms)
//
// Why is this necessary?
// if we want the method to work with other integer types, explicit type conversions would be necessary
// which in turn might cause problems. For example, uint allows you to represent values that are too big for an int.
//
// In short, type terms are used to specify operators (e.g. \ or %).
//
// That's why we need a way to specify that you can use / and %.
// In Go, this can be done with a 'type element' within an interface (e.g. see Integer below).
//
// Besides primitive predeclared types (such as int, float64, etc.),
// type terms can also be slices, maps, arrays, channels, structs, or even functions

// (1) this function only works with variables of type int
func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

// (2) fully functioning and defined generic function using a type element within the interface Integer
func divAndRemainderGeneric[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

// (3) fully functioning and defined generic function using a type element within the interface IntegerAllowsNonExactMatch
//
//	that allows using a user-defined type whose underlying type is one of the types listed in IntegerAllowsNonExactMatch
func divAndRemainderGenericAllowsNonExactMatch[T IntegerAllowsNonExactMatch](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

// NOTE: by default, 'type terms' match exactly
// if you try to use divAndRemainderGeneric() with a user-defined type whose underlying type
// is one of the types listed in Integer, you'll get an error
// (1)
type Integer interface {
	// ### 'type element' - start
	// each of the listed types is referred to as 'type term'
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
	// ### 'type element' - end

	// ### 'method element' - start
	//String() string
	// ### 'method element' - end
}

// (2)
// if you want a type term to be valid for any type that has the type term as its underlying type,
// put a tilde (~) before the type term
type IntegerAllowsNonExactMatch interface {
	// ### 'type element' - start
	// each of the listed types is referred to as 'type term'
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
	// ### 'type element' - end

	// ### 'method element' - start
	//String() string
	// ### 'method element' - end
}
