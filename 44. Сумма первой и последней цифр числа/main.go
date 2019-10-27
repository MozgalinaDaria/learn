package main

import "../base/console"

func main() {
	var firstFigure int

	number := console.ReadInt("Введите число: ")

	lastFigure := number % 10

	for  ; number / 10 != 0; number = number / 10 {
		firstFigure = number / 10
	}
	sum := firstFigure + lastFigure

	console.Writeln("Сумма первой и последней цифр: ",firstFigure , " + ", lastFigure, " = ", sum)
}
