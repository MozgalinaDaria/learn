package main

import "../base/console"

func main() {
	limit := console.ReadInt("Введите предел: ")
	sum := console.ReadInt("Введите искомую сумму: ")

	figuresBeforeLimit(limit, sum)
	figuresInLimit(limit, sum)
}

func figuresBeforeLimit(limit, sum int) {
	for i := 1; i < limit; i++ {
		secondFigure := sum - i - limit
		thirdFigure := limit

		for ; secondFigure <= limit; secondFigure++ {
			console.Writeln(i, secondFigure, thirdFigure)
			thirdFigure = thirdFigure - 1
		}
	}
}

func figuresInLimit(limit, sum int) {
	i := limit
	thirdFigure := limit - 1
	secondFigure := sum - i - thirdFigure

	for ; secondFigure < limit; secondFigure++ {
		console.Writeln(i, secondFigure, thirdFigure)
		thirdFigure = thirdFigure - 1
	}
}
