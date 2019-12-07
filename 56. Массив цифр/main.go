package main

import "../base/console"

func main() {
	var createdArray [] int

	number := console.ReadInt("Введите число: ")

	createdArray = getFigure(number)

	lengthOfArray := len(createdArray)
	console.Writeln(lengthOfArray)

	orderedArray(createdArray, lengthOfArray)
}

func getFigure(number int) [] int {
	var numberToArray int
	var createdArray [] int

	for ; number != 0; number = number / 10 {
		numberToArray = number % 10
		createdArray = append(createdArray, numberToArray)
	}
	return createdArray
}

func orderedArray(createdArray [] int, lengthOfArray int) {
	var orderedArray [] int

	index := lengthOfArray - 1

	for i := 0; i < lengthOfArray; i++ {
		orderedArray = append(orderedArray, createdArray[index])
		index -= 1
	}

	console.Writeln(orderedArray)
}
