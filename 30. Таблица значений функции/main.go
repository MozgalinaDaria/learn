package main

import (
	"../base/console"

)

func main(){
	var x, y float64

	console.Writeln("  x |  y  ")

	for x = -5; x <= 5; x += 0.5 {
    	y = 5 - x * x / 2
    	console.Writeln(int(x), " | ", y)
	}
}
