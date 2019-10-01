package main

import "../base/console"
import "../base/rand"
import "../base/math"

func main (){
	var random int
	random = rand.RandomInt(5, 15)
	console.Writeln("Рандомное число = ", random)

	console.Writeln("Числа от 0 до сгенерированного: ")
	for i := 0; i <= random; i++ {
		console.Write(i, " ")
	}

	console.Writeln()
	console.Writeln("Числа от сгенерированного до 0: ")

	for i := random; i >= 0; i-- {
		console.Write(i, " ")
	}

	if math.IsOdd(random) {
		console.Writeln()
		console.Writeln("Нечетные числа:")
		for i := 1; i <= random; i = i + 2 {
			console.Write(i, " ")
		}

	} else {
		console.Writeln()
		console.Writeln("Четные числа:")
		for i := 0; i <= random; i = i + 2 {
			console.Write(i, " ")
		}
	}
}

