package main

import (
	"../base/console"
)

func main() {
	var countOfRepeats int

	numberOfShops := console.ReadInt("Введите общее количество магазинов: ")
	namesOfShops := make(map[string]int)

	for i := 0; i < numberOfShops; i++ {
		namesOfShops[console.ReadString("Введите название магазина: ")]++
	}

	for _, value := range namesOfShops {
		if value > 1 {
			countOfRepeats += value - 1
		}
	}
	console.Writeln("Не посетили магазинов: ", countOfRepeats, "шт.")
}
