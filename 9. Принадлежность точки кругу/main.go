package main

import "../base/console"
import "math"

func main() {
	var centrX, centrY, radius, pointX, pointY, distanceBetweenPoint float64

	centrX = console.ReadFloat("Введите координату центра окружности по оси ОХ")
	centrY = console.ReadFloat("Введите координату центра окружности по оси ОY")
	radius = console.ReadFloat("Введите радиус")
	pointX = console.ReadFloat("Введите координаты точки по оси OX")
	pointY = console.ReadFloat("Введите координаты точки по оси OY")

	distanceBetweenPoint = getDistanceBetweenPoints(centrX, centrY, pointX, pointY)
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

func getDistanceBetweenPoints(x1, y1, x2, y2 float64) float64 {
	var distanceSquared float64

	distanceSquared = (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	return math.Pow(distanceSquared, 0.5)
}
