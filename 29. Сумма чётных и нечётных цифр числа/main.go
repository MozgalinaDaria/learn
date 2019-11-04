package main

import (
	"../base/console"
	"math"
)

func main() {
	var number, figure, sumOfEven, sumOfOdd int

	number = console.ReadInt("Введите число: ")
	for ; number != 0; number = number / 10 {
		figure = number % 10
		if figure%2 == 0 {
			sumOfEven = sumOfEven + figure
		} else {
			sumOfOdd = sumOfOdd + figure
		}
	}
	console.Writeln("Сумма чётных цифр числа = ", math.Abs(float64(sumOfEven)), ", сумма нечетных цифр числа = ", math.Abs(float64(sumOfOdd)))
}
