package main

import "fmt"

type Geometry interface {
	area() float64
	perim() float64
}

type Square struct {
	width  float64
	height float64
}

func newSquare(width float64, height float64) Square {
	return Square{
		width,
		height,
	}
}

func (s *Square) setSquare(width float64, height float64) {
	s.width = width
	s.height = height
}

func main() {
	g := newSquare(3, 4)
	fmt.Printf("%+v\n", g)
}
