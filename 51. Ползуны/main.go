package main

import "../base/console"

func main() {
	var totalNumberofPairsOfWheels, diameter int

	diameterOfWheels := make(map[int]int)
	countOfPairsOfWheels := console.ReadInt("Введите общее количество колесных пар: ")

	for i := 0; i < countOfPairsOfWheels; i++ {
		diameter = console.ReadInt("Введите диаметр колеса: ")
		diameterOfWheels[diameter]++
	}

	for _, value := range diameterOfWheels {
		if value%4 == 0 {
			totalNumberofPairsOfWheels++
		}
	}
	console.Writeln("Всего можно будет оснастить вагонов: ", totalNumberofPairsOfWheels)
}
