package main

import (
	"errors"
	"fmt"
)

func main() {

	var err error

	err = funcThatReturnsAnError()

	// to handle different error types (wrapped and non-wrapped ones)
	// you can use a type switch
	switch err := err.(type) {
	case interface{ Unwrap() error }:
		innerErr := err.Unwrap()
		fmt.Println(innerErr)
	case interface{ Unwrap() []error }:
		innerErrs := err.Unwrap()
		for _, err := range innerErrs {
			fmt.Println(err)
		}
	default:
		fmt.Println(err)
	}
}

func funcThatReturnsAnError() error {
	return errors.New("something bad happened")
}

type MyError struct {
	Code   int
	Errors []error
}

func (m MyError) Error() string {
	return errors.Join(m.Errors...).Error()
}

func (m MyError) Unwrap() []error {
	return m.Errors
}
