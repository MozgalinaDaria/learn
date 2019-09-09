package main

import "../dexes/console"
import "../dexes/rand"

func main () {
	console.Write("Рандомное трехзначное число = ", randomNumber())
}

func randomNumber() int {
	var threeDigitNumber int

	threeDigitNumber = rand.RandomInt(100, 999)
	return threeDigitNumber - threeDigitNumber % 10
}
