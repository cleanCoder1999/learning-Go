package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ### - exercise 1: generic function that doubles each integer and float type
	{
		fmt.Println(double(uint8(34)))
		fmt.Println(double(uint64(1)))
		fmt.Println(double(int32(1)))
		fmt.Println(double(int64(33)))
		fmt.Println(double(int64(7878)))
		fmt.Println(double(float32(7878)))
		fmt.Println(double(float64(7878)))
		fmt.Println("")
	}

	// ### - exercise 2: generic interface Printable
	{
		var i MyGenericInt = 123
		printPrintable(i)
		var f MyGenericFloat = 123.34
		printPrintable(f)
		fmt.Println("")
	}

	// ### - exercise 3:
	{
		l := SinglyLinkedNode[int]{
			val: 123,
		}
		l.Add(100)
		l.Add(200)
		l.Add(300)
		l.Add(400)
		l.Add(500)
		i := l.Index(300)
		fmt.Println("index of 300 is:", i)
		fmt.Println("")
		l.Insert(1500, 3)
		l.Insert(90, 0)

		for l.next != nil {
			fmt.Println("val:", l.val)
			l = *l.next
		}
	}
}

// ### - exercise 1
func double[T Number](v T) T {
	return v * 2
}

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// ### - exercise 2
type MyGenericInt int

func (m MyGenericInt) String() string {
	return "MyGenericType: " + strconv.Itoa(int(m))
}

type MyGenericFloat float64

func (m MyGenericFloat) String() string {
	return "MyGenericType: " + fmt.Sprintf("%v", float64(m))
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

func printPrintable[T Printable](v T) {
	fmt.Println(v)
}

// ### - exercise 3
type SinglyLinkedNode[T comparable] struct {
	val  T
	next *SinglyLinkedNode[T]
}

func (l *SinglyLinkedNode[T]) Add(val T) {

	if l.next == nil {
		l.next = &SinglyLinkedNode[T]{
			val: val,
		}
		return
	}

	l.next.Add(val)
}

func (l *SinglyLinkedNode[T]) Index(v T) int {
	return l.indexHelper(v, 0)
}

func (l *SinglyLinkedNode[T]) indexHelper(v T, i int) int {

	if l.val == v {
		return i
	}

	if l.next == nil {
		return -1
	}

	i++
	return l.next.indexHelper(v, i)
}

func (l *SinglyLinkedNode[T]) Insert(v T, i int) bool {
	if i < 0 {
		return false
	}

	// given index is greater than length of list
	if l.next == nil && i > 1 {
		return false
	}

	if i > 0 {
		i--
		return l.next.Insert(v, i)
	}

	toBeMoved := SinglyLinkedNode[T]{
		val:  l.val,
		next: l.next,
	}

	l.val = v
	l.next = &toBeMoved

	return true
}
