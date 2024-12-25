package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	// ### - simple calculator: function map
	{
		expressions := [][]string{
			{"2", "+", "3"},
			{"2", "-", "3"},
			{"2", "*", "3"},
			{"2", "/", "3"},
			{"2", "/", "0"},
			{"2", "%", "3"},
			{"two", "+", "three"},
			{"5"},
		}

		for _, expression := range expressions {
			if len(expression) != 3 {
				fmt.Println("Invalid expression:", expression)
				continue
			}

			p1, err := strconv.Atoi(expression[0])
			if err != nil {
				fmt.Println(err)
				continue
			}

			op := expression[1]
			opFunc, ok := opMap[op]
			if !ok {
				fmt.Println("Unsupported operator:", p1)
				continue
			}

			p2, err := strconv.Atoi(expression[2])
			if err != nil {
				fmt.Println(err)
				continue
			}

			result, err := opFunc(p1, p2)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(result)
		}
		fmt.Println("")
	}

	// ### - fileLen: defer with closure
	{
		fileName := "dummy.txt"
		fl, err := fileLen(fileName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("file length:", fl)
		fmt.Println("")
	}

	// ### - prefixer: function return value
	{
		helloPrefix := prefixer("Hello")
		fmt.Println(helloPrefix("Bob"))
		fmt.Println(helloPrefix("Maria"))
	}
}

var opMap = map[string]func(a, b int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func add(a, b int) (int, error) { return a + b, nil }

func sub(a, b int) (int, error) { return a - b, nil }

func mul(a, b int) (int, error) { return a * b, nil }

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	var total int
	var count int
	buf := make([]byte, 1024)

	for {
		count, err = f.Read(buf)
		total += count

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	return total, nil
}

func prefixer(prefix string) func(string) string {
	return func(s string) string { return prefix + " " + s }
}
