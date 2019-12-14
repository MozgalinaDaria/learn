package main

import "../base/console"

const Square = 1
const Rectangle = 2
const Triangle = 3
const Circle = 4
const Pi = 3.14

func main() {
	figure := console.ReadFloat(
		"Введите номер фигуры: ",
		Square, " - квадрат, ",
		Rectangle, " - прямоугольник, ",
		Triangle, " - треугольник, ",
		Circle, " - круг",
	)

	console.Writeln("Площадь фигуры = ", GetArea(figure))
}

func GetArea(figure float64) float64 {
	switch figure {
	case Square:
		GetSquareArea()
	case Rectangle:
		GetRectangleArea()
	case Triangle:
		GetTriangleArea()
	case Circle:
		GetCircleArea()
	}

	panic("Неизвестная фигура")
}

func GetSquareArea() float64 {
	length := console.ReadFloat("Введите длину стороны: ")

	return length * length
}

func GetRectangleArea() float64 {
	length := console.ReadFloat("Введите длину: ")
	width := console.ReadFloat("Введите ширину: ")

	return length * width
}

func GetTriangleArea() float64 {
	length := console.ReadFloat("Введите длину основания: ")
	height := console.ReadFloat("Введите высоту: ")

	return 0.5 * length * height
}

func GetCircleArea() float64 {
	r := console.ReadFloat("Введите длину радиуса: ")

	return Pi * r * r
}
