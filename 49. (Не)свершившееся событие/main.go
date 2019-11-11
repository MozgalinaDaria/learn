package main

import (
	"../base/console"
	"../base/rand"
)

func main() {
	var randomNumber float64

	a := console.ReadFloat("Введите вероятность свершения события от 0 до 1: ")
	probability := a * 100

	for i := 0; i < 100; i++ {
		randomNumber = float64(rand.RandomInt(0, 99))

		if probability >= randomNumber {
			console.Writeln("Событие свершилось")
		} else {
			console.Writeln("Событие не свершилось")
		}
	}
}
