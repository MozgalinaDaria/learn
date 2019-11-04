package main

import (
	"../base/console"
	"math"
)

func main() {
	var newNumber, AbsolutValueOfDegree int

	number := console.ReadInt("Назовите число: ")
	degree := console.ReadInt("Назовите степень, в которую нужно возвести ", number, ": ")

	if degree == 0 {
		console.Writeln(number, " в степени ", degree, "равен 1")
		return
	}
	if degree != 0 {
		console.Write(number, " в степени ", degree, " = ")

		newNumber = number
		AbsolutValueOfDegree = int(math.Abs(float64(degree)))

		for i := 0; i <= AbsolutValueOfDegree-2; i++ {
			newNumber *= number
		}
		if degree > 0 {
			console.Write(newNumber)
		}
		if degree < 0 {
			console.Write("1/", newNumber)
		}
	}
}
