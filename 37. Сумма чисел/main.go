package main

import (
"../base/console"
"../base/rand"
)

func main() {
	var sumFromTenCycle, sumFromTenProgression, amountOfNumbersInProgression int

	number := rand.RandomInt(50, 100)
	console.Writeln("Рандомно было сгенерировано число: ", number)

	for i := 10; i <= number; i++ {

		sumFromTenCycle = sumFromTenCycle + i
	}
	console.Writeln("Сумма всех чисел от 10 до ", number, " = ", sumFromTenCycle)

	amountOfNumbersInProgression = (number - 10 + 1)
	sumFromTenProgression = (10 + number) * amountOfNumbersInProgression / 2
	console.Writeln("Посчитано по формуле арифметической прогрессии. Сумма всех чисел от 10 до ", number, " = ", sumFromTenProgression)
}

