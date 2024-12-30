package main

import (
	"fmt"
	"math"
)

func main() {
	pair2Da := Pair[Point2D]{Point2D{1, 7}, Point2D{3, 4}}
	pair2Db := Pair[Point2D]{Point2D{88, 9}, Point2D{3, 7}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)

	pair2Da = Pair[Point2D]{Point2D{1, 2}, Point2D{3, 4}}
	pair2Db = Pair[Point2D]{Point2D{5, 9}, Point2D{3, 7}}
	closer = FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)

	pair3Da := Pair[Point3D]{Point3D{1, 2, 5}, Point3D{5, 4, 9}}
	pair3Db := Pair[Point3D]{Point3D{5, 9, 5}, Point3D{3, 7, 9}}
	closer2 := FindCloser(pair3Da, pair3Db)
	fmt.Println(closer2)

}

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)

	if d1 < d2 {
		return pair1
	}

	if d1 == d2 {
		fmt.Println("both pairs are equally close")
	}

	return pair2
}

// ### Point2D
type Point2D struct {
	x, y int
}

func (p2d Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2d.x, p2d.y)
}

func (p2d Point2D) Diff(d2 Point2D) float64 {
	x := p2d.x - d2.x
	y := p2d.y - d2.y

	return math.Sqrt(float64(x*x + y*y))
}

// ### Point3D
type Point3D struct {
	x, y, z int
}

func (p3d Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p3d.x, p3d.y, p3d.z)
}

func (p3d Point3D) Diff(d3 Point3D) float64 {
	x := p3d.x - d3.x
	y := p3d.y - d3.y
	z := p3d.z - d3.z

	return math.Sqrt(float64(x*x + y*y + z*z))
}
