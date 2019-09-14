package main

import "../dexes/console"

func main() {
	var number int

	console.Writeln("Введите номер фигуры (1 - квадрат, 2 - прямоугольник, 3 - треугольник, 4 - круг)")
	number = console.ReadInt()

	console.Writeln("Вами выбран ", number)
	switch number {
	case 1:
		console.Writeln("Квадрат")
	case 2:
		console.Writeln("Прямоугольник")
	case 3:
		console.Writeln("Треугольник")
	case 4:
		console.Writeln("Круг")
	}
}