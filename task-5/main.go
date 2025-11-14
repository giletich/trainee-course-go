package main

import (
	"flag"
	"fmt"
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
	fmt.Println(c.Radius * 2 * 3.14)
}

func (r Rectangle) Area() {
	fmt.Println(r.Height*r.Width)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	shape := flag.String("shape", "", "youre figure")
	radius := flag.Float64("radius", 0, "radius of circle")
	width := flag.Float64("width", 0, "width of rectangle")
	height := flag.Float64("height", 0, "height of rectangle")

	flag.Parse()

	var area Shape

	if *shape != "circle" && *shape != "rectangle" {
		fmt.Println("set up shape")
		return
	}
	if *radius < 0 || *width < 0 || *height < 0 {
		fmt.Println("set up non negative value")
		return
	}

	if *shape == "circle" {
		if *radius == 0 {
			fmt.Println("set up radius")
			return
		}
		area = Circle{Radius: *radius}
		CountArea(area)
		return
	}
	if *shape == "rectangle" {
		if *width == 0 || *height == 0 {
			fmt.Println("set up width or height")
			return
		}
		area = Rectangle{Width: *width, Height: *height}
		CountArea(area)
		return
	}
}