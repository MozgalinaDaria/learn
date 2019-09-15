package main

import "../base/console"
import "math"

func main() {
	var centrX, centrY, radius, pointX, pointY, distanceBetweenPoint float64

	console.Writeln("Введите координаты центра окружности")
	centrX = console.ReadFloat()
	centrY = console.ReadFloat()
	console.Writeln("Введите радиус")
	radius = console.ReadFloat()
	console.Writeln("Введите координаты точки")
	pointX = console.ReadFloat()
	pointY = console.ReadFloat()

	distanceBetweenPoint = getDistance(centrX, centrY, pointX, pointY)
	if distanceBetweenPoint < radius {
		console.Writeln("Точка лежит внутри окружности")
		return
	}

	if distanceBetweenPoint == radius {
		console.Writeln("Точка лежит на границе окружности")
		return
	}

	console.Writeln("Точка не принадлежит окружности")
}

func getDistance(x1, y1, x2, y2 float64) float64 {
	var distanceSquared float64

	distanceSquared = (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	return math.Pow(distanceSquared, 0.5)
}
