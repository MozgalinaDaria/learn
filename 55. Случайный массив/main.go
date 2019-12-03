package main

import (
	"../base/console"
	"../base/rand"
)

func main() {
	var randomArray [] int
	var randomNumber int

	lengthOfArray := console.ReadInt("Введите длину массива: ")

	for i := 0; i < lengthOfArray; i++ {
		randomNumber = rand.RandomInt(i, 10000)
		randomArray = append(randomArray, randomNumber)
	}

	console.Writeln("Массив рандомных чисел: ", randomArray)
}
