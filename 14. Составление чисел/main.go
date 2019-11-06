package main

import "../base/console"

func main() {

	number := console.ReadInt("Введите любое трехзначное число: ")

	figure1, figure2, figure3 := GetFigure(number)

	console.Writeln("Трехзначные числа, которые можно получить из цифр введенного числа: ")
	console.Writeln(figure1, figure2, figure3)
	console.Writeln(figure1, figure3, figure2)
	console.Writeln(figure2, figure1, figure3)
	console.Writeln(figure2, figure3, figure1)
	console.Writeln(figure3, figure1, figure2)
	console.Writeln(figure3, figure2, figure1)
}

func GetFigure(number int) (int, int, int) {

	figure3 := number % 10
	figure2 := (number % 100) / 10
	figure1 := number / 100

	return figure3, figure2, figure1
}
