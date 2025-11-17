package main

import (
	"flag"
	"fmt"
	"task-5/shapes"
)

var (
	circleShape = "circle"
	rectangleShape = "rectangle"
)


func main() {
	
	shape := flag.String("shape", "", "youre figure")
	radius := flag.Float64("radius", 0, "radius of circle")
	width := flag.Float64("width", 0, "width of rectangle")
	height := flag.Float64("height", 0, "height of rectangle")

	flag.Parse()

	var area shapes.Shape

	if *shape != circleShape && *shape != rectangleShape {
		fmt.Println("set up shape")
		return
	}

	if *radius < 0 || *width < 0 || *height < 0 {
		fmt.Println("set up non negative value")
		return
	}

	if *shape == circleShape {

		if *radius == 0 {
			fmt.Println("set up radius")
			return
		}

		area = shapes.Circle{Radius: *radius}
		shapes.CountArea(area)
		return
	}
	if *shape == rectangleShape {

		if *width == 0 || *height == 0 {
			fmt.Println("set up width or height")
			return
		}

		area = shapes.Rectangle{Width: *width, Height: *height}
		shapes.CountArea(area)
		return
	}
}