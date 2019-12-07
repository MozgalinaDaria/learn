package main

import "../base/console"

func main() {
	var c, r Shape
	c = 1
	r = 2
	console.Writeln(totalArea(&c, &r))

}

type Shape interface {
	Area() float64
}

func totalArea(shapes ...*Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}
