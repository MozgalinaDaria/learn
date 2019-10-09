package main

import (
	"../base/console"
	"math"
)

func main() {
	var number, numberForCycle, figure int

	number = console.ReadInt("Введите число: ")
	numberForCycle = int(math.Abs(float64(number)))
	if number >= 0 {
		console.Write("Перевернутое число: ")
	}
	if number < 0 {
		console.Write("Перевернутое число: -")
	}

	for ; numberForCycle != 0; numberForCycle = numberForCycle / 10 {
		figure = numberForCycle % 10
		console.Write(figure)
	}
}
