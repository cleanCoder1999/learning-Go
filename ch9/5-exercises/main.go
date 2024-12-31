package main

import (
	"errors"
	"fmt"
)

var (
	ErrMySentinel     = func() error { return errors.New("my sentinel; stop immediately") }()
	ErrSecondSentinel = func() error { return errors.New("second sentinel; jeiij") }()
)

func main() {

	// ### - exercise 1: create custom Sentinel error
	{
		err := myFunc()
		if errors.Is(err, ErrMySentinel) {
			fmt.Println("a sentinel error occurred")
		}
		fmt.Println("")
	}

	// ### - exercise 2: define a custom error type
	{
		var emptyEmployeeErr EmptyEmployeeErr
		e := Employee{FirstName: "Dan"}

		err := validateEmployee(e)
		if errors.As(err, &emptyEmployeeErr) {
			fmt.Println(emptyEmployeeErr)
		}
		fmt.Println("")
	}

	// ### - exercise 3: wrap all errors in a single error
	{
		empl := Employee{}
		err := validateEmployee(empl)

		var validationErrors ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors.Unwrap() {
				fmt.Println(e)
			}
		}
	}
}

func myFunc() error {
	return ErrMySentinel
}

// NOTE: always return the interface 'error' even if your function returns an instance of a custom error,
//       so that third-parties do not have to rely on your custom errors

// this differs from the general rule: "accept interfaces, return structs"
func validateEmployee(e Employee) error {

	errs := ValidationErrors{
		Errors: []error{},
	}

	if e.FirstName == "" && e.LastName == "" {
		errs.Errors = append(errs.Errors, errors.New("no names provided"))
	}

	if e.Age == 0 {
		errs.Errors = append(errs.Errors, errors.New("no age provided"))
	}

	if e.FirstName == "" || e.LastName == "" || e.Age == 0 {
		errs.Errors = append(errs.Errors, EmptyEmployeeErr{e.FirstName})
	}

	return errs
}

type ValidationErrors struct {
	Errors []error
}

func (v ValidationErrors) Error() string {
	return errors.Join(v.Errors...).Error()
}

func (v ValidationErrors) Unwrap() []error {
	return v.Errors
}

type EmptyEmployeeErr struct {
	Name string
}

func (e EmptyEmployeeErr) Error() string {
	return fmt.Sprintf("employee %s is empty", e.Name)
}

func (e EmptyEmployeeErr) Unwrap() error {
	return e
}

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}