package main

import "../base/console"
import "math"

func main() {
	var a, b, c, doubleA, negativeB, discriminant, rootOfDiscriminant float64
	var xFirst, xSecond float64

	console.Write("Введите a: ")
	a = console.ReadFloat()
	console.Write("Введите b: ")
	b = console.ReadFloat()
	console.Write("Введите c: ")
	c = console.ReadFloat()

	discriminant = Discriminant(a, b, c)
	if discriminant < 0 {
		console.Write("Данное уравнение корней не имеет")
		return
	}

	doubleA = 2 * a
	rootOfDiscriminant = math.Sqrt(discriminant)
	negativeB = -b

	xFirst = (negativeB + rootOfDiscriminant) / doubleA

	if discriminant == 0 {
		console.Write("Корень уравнения = ", xFirst)
		return
	}

	xSecond = (negativeB - rootOfDiscriminant) / doubleA
	console.Write("Первый корень = ", xFirst, " ", "Второй корень = ", xSecond)
}

func Discriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}
