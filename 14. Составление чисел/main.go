package main

import "../dexes/console"

func main () {
	var number, figure3, figure2, figure1 int
	console.Writeln("Введите любое трехзначное число")
	number = console.ReadInt()

	figure1, figure2, figure3 = GetFigure (number)

	console.Writeln(figure1, figure2, figure3)
	console.Writeln(figure1, figure3, figure2)
	console.Writeln(figure2, figure1, figure3)
	console.Writeln(figure2, figure3, figure1)
	console.Writeln(figure3, figure1, figure2)
	console.Writeln(figure3, figure2, figure1)
}

func GetFigure (number int) (int, int, int) {
	var figure1, figure2, figure3 int
	figure3 = number % 10
	figure2 = (number % 100) / 10
	figure1 = number / 100

	return figure3, figure2, figure1
}



