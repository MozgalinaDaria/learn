package main

import "../base/console"

const Square = 1
const Rectangle = 2
const Triangle = 3
const Circle = 4
const Pi = 3.14

func main() {
	var figure float64

	console.Writeln("Введите номер фигуры: ", Square, " - квадрат, ", Rectangle, " - прямоугольник, ",
		Triangle, " - треугольник, ", Circle, " - круг")
	figure = console.ReadFloat()

	console.Writeln("Площадь фигуры = ", GetSquare(figure))
}

func GetSquare(figure float64) float64 {
	// TODO вынести расчёт площади для каждой фигуры в отдельную функцию

	switch figure {
	case Square:
		var length float64

		console.Writeln("Введите длину стороны")
		length = console.ReadFloat()
		return length * length

	case Rectangle:
		var length, width float64

		console.Writeln("Введите длину и ширину")
		length = console.ReadFloat()
		width = console.ReadFloat()
		return length * width

	case Triangle:
		var length, height float64

		console.Writeln("Введите длину основания и высоту")
		length = console.ReadFloat()
		height = console.ReadFloat()
		return 0.5 * length * height

	case Circle:
		var r float64

		console.Writeln("Введите длину радиуса")
		r = console.ReadFloat()
		return Pi * r * r
	}

	panic("Неизвестная фигура")
}
