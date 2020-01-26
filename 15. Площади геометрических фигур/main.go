package main

import "../base/console"

const SquareNumber = 1
const RectangleNumber = 2
const TriangleNumber = 3
const CircleNumber = 4
const Pi = 3.14

type Square struct {
	Length float64
}

func (object *Square) GetArea() float64 {

	return object.Length * object.Length
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (object *Rectangle) GetArea() float64 {

	return object.Height * object.Width
}

type Triangle struct {
	Length float64
	Height float64
}

func (object *Triangle) GetArea() float64 {

	return 0.5 * object.Length * object.Height
}

type Circle struct {
	Radius float64
}

func (object *Circle) GetArea() float64 {

	return Pi * object.Radius * object.Radius
}

type Shape interface {
	GetArea() float64
}

func MakeShape(figureNumber int) Shape {
	switch figureNumber {
	case SquareNumber:
		length := console.ReadFloat("Введите длину стороны квадрата: ")
		return &Square{Length: length}

	case RectangleNumber:
		height := console.ReadFloat("Введите высоту: ")
		width := console.ReadFloat("Введите ширину: ")
		return &Rectangle{Height: height, Width: width}

	case TriangleNumber:
		length := console.ReadFloat("Введите длину основания: ")
		height := console.ReadFloat("Введите высоту: ")
		return &Triangle{Length:length, Height:height}

	case CircleNumber:
		r := console.ReadFloat("Введите длину радиуса: ")
		return &Circle{Radius:r}
	}

	panic("Неизвестная фигура")
}

func main() {
	figureNumber := console.ReadInt(
		"Введите номер фигуры: ",
		SquareNumber, " - квадрат, ",
		RectangleNumber, " - прямоугольник, ",
		TriangleNumber, " - треугольник, ",
		CircleNumber, " - круг",
	)

	figure := MakeShape(figureNumber)

	console.Writeln("Площадь фигуры = ", figure.GetArea())
}
