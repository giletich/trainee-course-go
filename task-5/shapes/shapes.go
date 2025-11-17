package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area()
}

func CountArea(s Shape) {
	s.Area()
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() {
	fmt.Println(c.Radius * 2 * math.Pi)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() {
	fmt.Println(r.Height * r.Width)
}
