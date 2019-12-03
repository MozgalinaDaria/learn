package main

import "../base/console"

func main() {
	var numberToArray int
	var createdArray [] int
	var serialNumber int

	number := console.ReadInt("Введите число: ")

	for ; number != 0; number = number / 10 {
		numberToArray = number % 10
		createdArray = append(createdArray, numberToArray)
	}

	lengthOfArray := len(createdArray)
	for i:= 0; i < lengthOfArray; i++ {

	}
}
