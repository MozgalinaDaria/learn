package main

import (
	"../base/console"
	"sort"
)

func main() {
	var figures []int
	var figure, lenOfArray int
	number := console.ReadInt("Введите число: ")

	for ; number > 0; number = number / 10 {
		figure = number % 10
		figures = append(figures, figure)
	}

	lenOfArray = len(figures) - 1

	sort.Ints(figures)

	if figures [0] == figures [lenOfArray] {
		console.Writeln("В числе все цифры равны ", figures [0])
	} else {
		console.Writeln("Минимальная цифра = ", figures [0], " , максимальная цифра равна = ", figures [lenOfArray])
	}
}
