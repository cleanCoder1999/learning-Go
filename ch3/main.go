package main

import "fmt"

func main() {
	// ### confusing slices
	{
		x := make([]string, 0, 5)
		fmt.Println("x:", x)
		x = append(x, "a", "b", "c", "d")
		fmt.Println("x:", x)

		// full slice expression protects against confusing behavior due to append and memory sharing
		// s[<beginning (incl)> : <end (excl)> : <last position in the parent slice's capacity>]
		// sub-slice capacity = last position in the parent slice's capacity - beginning (incl)
		y := x[:2:2]
		z := x[2:4:4]
		fmt.Println(cap(x), cap(y), cap(z))
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)

		fmt.Println("append to Y: i,j,k")
		y = append(y, "i", "j", "k")

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)

		fmt.Println("append to X: x")
		x = append(x, "x")

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)

		fmt.Println("append to Z: y")
		z = append(z, "y")

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		fmt.Println("")
		fmt.Println("")
	}

	// ### copying slices
	{
		a := []int{1, 2, 3, 4}
		b := make([]int, 4)
		num := copy(b, a)
		fmt.Println(b, num)
		fmt.Println("")
	}

	// ### converting strings to rune and byte slices
	{
		var s = "Hello, ☀️"
		bs := []byte(s)
		rs := []rune(s)
		fmt.Println(bs)
		fmt.Println(rs)
		fmt.Println("")
	}

	// ### using a map as set
	{
		intSet := map[int]bool{}
		vals := []int{1, 2, 3, 4, 5, 5, 6, 8, 8, 8, 8, 8, 8, 7, 8, 9, 0, 12, 55}

		// add the values to the map
		for _, v := range vals {
			intSet[v] = true
		}

		fmt.Println("len(vals): ", len(vals))
		fmt.Println("len(intSet): ", len(intSet))

		fmt.Println(intSet[5])
		fmt.Println(intSet[500])

		if intSet[55] {
			fmt.Println("55 is in the set")
		}
		fmt.Println("")
	}

	// ### exercise 1
	{
		greetings := []string{"Hello", "Hola", "Привет", "안녕하세요", "こんにちは"}
		sub1 := greetings[:2:3]
		sub2 := greetings[1:4:4]
		fmt.Println(sub1)
		fmt.Println(sub2)
		fmt.Println("")
	}

	// ### exercise 2 - print the fourth rune of message as a character, not a number
	{
		message := "Hi 👩🏻 and 👨🏻"
		runeMsg := []rune(message)

		fmt.Println(runeMsg)
		fmt.Println("# of bytes:", len(string(runeMsg[3])))
		fmt.Println(string(runeMsg[3]))
		fmt.Printf("%c\n", runeMsg[3])
		fmt.Println("")
	}

	// ### exercise 2 - print the fourth rune of message as a character, not a number
	{
		type Employee struct {
			firstName string
			lastName  string
			id        int
		}

		ken := Employee{
			"Ken",
			"Thompson",
			1,
		}

		dennis := Employee{
			firstName: "Dennis",
			lastName:  "Ritchie",
			id:        2,
		}

		var rob Employee
		rob.firstName = "Rob"
		rob.lastName = "Pike"
		rob.id = 3

		fmt.Println(ken)
		fmt.Println(dennis)
		fmt.Println(rob)

		fmt.Println("")
	}

}
