package main

import (
	"fmt"
	"math"
)

func main() {
	a := NewPoint(3, 2)
	b := NewPoint(7, 8)

	fmt.Println(Distance(a, b))
}

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (point Point) X() float64 {
	return point.x
}

func (point Point) Y() float64 {
	return point.y
}

func Distance(a, b Point) float64 {
	dx, dy := Vector(a, b)
	return math.Sqrt(math.Pow(dx, 2.0) + math.Pow(dy, 2.0))
}

func Vector(a, b Point) (float64, float64) {
	return a.X() - b.X(), a.Y() - b.Y()
}
