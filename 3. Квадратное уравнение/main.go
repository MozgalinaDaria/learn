package main

import "../base/console"
import "math"

func main() {

	a := console.ReadFloat("Введите a: ")
	b := console.ReadFloat("Введите b: ")
	c := console.ReadFloat("Введите c: ")

	discriminant := Discriminant(a, b, c)
	if discriminant < 0 {
		console.Write("Данное уравнение корней не имеет")
		return
	}

	doubleA := 2 * a
	rootOfDiscriminant := math.Sqrt(discriminant)
	negativeB := -b

	xFirst := (negativeB + rootOfDiscriminant) / doubleA

	if discriminant == 0 {
		console.Write("Корень уравнения = ", xFirst)
		return
	}

	xSecond := (negativeB - rootOfDiscriminant) / doubleA
	console.Write("Первый корень = ", xFirst, " ", "Второй корень = ", xSecond)
}

func Discriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}
