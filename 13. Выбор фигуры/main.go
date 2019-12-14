package main

import "../base/console"

const Square = 1
const Rectangle = 2
const Triangle = 3
const Circle = 4

func main() {
	number := console.ReadInt(
		"Введите номер фигуры: ",
		Square, " - квадрат, ",
		Rectangle, " - прямоугольник, ",
		Triangle, " - треугольник, ",
		Circle, " - круг",
	)

	console.Write("Вами выбран ")
	switch number {
	case Square:
		console.Write("квадрат")
	case Rectangle:
		console.Write("прямоугольник")
	case Triangle:
		console.Write("треугольник")
	case Circle:
		console.Write("круг")
	}
}
