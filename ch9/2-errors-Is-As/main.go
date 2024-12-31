package main

import (
	"errors"
	"fmt"
	"os"
)

type MyErr struct {
	Codes []int
}

func main() {
	// ### errors.Is()
	{
		err := fileChecker("test.txt")
		if err != nil {

			// errors.Is() returns true if any error in the error tree matches the Sentinel error os.ErrNotExist
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println("file does not exist")
			}
		}
	}

	// ### errors.As()
	{
		err := funcThatReturnsAnError()
		var myErr MyErr
		if errors.As(err, &myErr) {
			fmt.Println(myErr.Codes)
		}
	}
}

func funcThatReturnsAnError() error {
	return errors.New("something bad happened")
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}
