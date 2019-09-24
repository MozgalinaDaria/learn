package main

import (
	"../base/console"
	"math"
)

type Rectangle struct{
	Left float64
	Right float64
	Top float64
	Bottom float64
}

func main(){
	var r1, r2 Rectangle

	r1.Left = console.ReadFloat("Введите x верхнего левого угла № 1: ")
	r1.Top = console.ReadFloat("Введите y верхнего левого угла № 1: ")
	r1.Right = console.ReadFloat("Введите x верхнего правого угла № 1: ")
	r1.Bottom = console.ReadFloat("Введите y нижнего правого угла № 1: ")

	r2.Left = console.ReadFloat("Введите x верхнего левого угла № 2: ")
	r2.Top = console.ReadFloat("Введите y верхнего левого угла № 2: ")
	r2.Right = console.ReadFloat("Введите x верхнего правого угла № 2: ")
	r2.Bottom = console.ReadFloat("Введите y нижнего правого угла № 2: ")

	if AreCrossing (r1, r2) {
		console.Writeln("Прямоугольники пересекаются")

		console.Writeln("Площадь пересечения равна ", IntersectionSquare (r1, r2))

		if AreNested(r1, r2) == 0 {
			console.Writeln("Прямоугольник 1 вложен в прямоугольник 2")
			return
		}

		if AreNested(r1, r2) == 1 {
			console.Writeln("Прямоугольник 2 вложен в прямоугольник 1")
			return
		}

		if AreNested(r1, r2) == -1 {
			console.Writeln("Прямоугольники не вложены друг в друга ")
			return
		}

		return
	} else {console.Writeln("Прямоугольники не пересекаются")}

}

func AreCrossing(r1, r2 Rectangle) bool {
	if (r1.Top < r2.Bottom || r2.Bottom < r1.Bottom) && (r1.Left < r2.Right || r2.Right < r1.Right){
		return true
	}

	if (r2.Top < r1.Bottom || r1.Bottom < r2.Bottom) && (r2.Left < r1.Right || r1.Right < r2.Right){
		return true
	}

	return false
}

func IntersectionSquare(r1, r2 Rectangle) float64 {
	var OX, OY float64

	OX = math.Max(r1.Left, r2.Left) - math.Min(r1.Right, r2.Right)
	OY = math.Min(r1.Top, r2.Top) - math.Max(r1.Bottom, r2.Bottom)

	return math.Abs(OX * OY)
}

func AreNested(r1, r2 Rectangle) int {
	if r1.Top < r2.Top && r1.Bottom > r2.Bottom && r1.Left > r2.Left && r1.Right < r2.Right {
		return 0
	}

	if r2.Top < r1.Top && r2.Bottom > r1.Bottom && r2.Left > r1.Left && r2.Right < r1.Right {
		return 1
	} else {
		return -1
	}
}

