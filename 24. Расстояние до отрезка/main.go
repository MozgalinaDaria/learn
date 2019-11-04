package main

import (
	"../base/console"
	"math"
)

const PointIsInSegment = 0
const PointIsToTheLeftOfSegment = 1
const PointIsToTheRightOfSegment = -1

func main() {
	var xFirst, хSecond, yFirst, ySecond, x, y float64
	var occasion int

	xFirst = console.ReadFloat("Введите xFirst отрезка: ")
	yFirst = console.ReadFloat("Введите yFirst отрезка: ")
	хSecond = console.ReadFloat("Введите хSecond отрезка: ")
	ySecond = console.ReadFloat("Введите ySecond отрезка: ")
	x = console.ReadFloat("Введите х точки: ")
	y = console.ReadFloat("Введите y точки: ")

	occasion = IsInSegment(xFirst, хSecond, yFirst, ySecond, x, y)
	console.Writeln(occasion)
	console.Writeln("Расстояние от точки до отрезка составляет: ", GetDistance(occasion, xFirst, хSecond, yFirst, ySecond, x, y))
}

func IsInSegment(xFirst, хSecond, yFirst, ySecond, x, y float64) int {
	if xFirst <= x && x <= хSecond || yFirst <= y && y <= ySecond {

		return PointIsInSegment
	}

	if x < xFirst || y < ySecond {

		return PointIsToTheLeftOfSegment
	}

	return PointIsToTheRightOfSegment
}

func GetDistance(a int, хFirst, хSecond, yFirst, ySecond, x, y float64) float64 {
	switch a {
	case PointIsInSegment:
		return math.Sqrt((хFirst-x)*(хFirst-x) + (ySecond-y)*(ySecond-y))
	case PointIsToTheLeftOfSegment:
		return math.Sqrt((хFirst-x)*(хFirst-x) + (ySecond-y)*(ySecond-y))
	default:
		return math.Sqrt((хSecond-x)*(хSecond-x) + (yFirst-y)*(yFirst-y))
	}
}
