package main

import (
	"../base/console"
	"../base/rand"
	"time"
)

func main() {
	var i, j, sumFromTenCycle, sumFromTenProgression, amountOfNumbersInProgression int

	number := rand.RandomInt(50, 100)
	console.Writeln("Рандомно было сгенерировано число: ", number)

	var startTimeCycle = time.Now()

	for i = 0; i < 100000000; i++ {
		for j = 10; j <= number; j++ {
			sumFromTenCycle += j
		}
	}
	var endTimeCycle = time.Since(startTimeCycle)

	console.Writeln("Посчитано в цикле. Сумма всех чисел от 10 до ", number, " = ", sumFromTenCycle/i,
		"; время обработки цикла, сек = ", endTimeCycle.Seconds())

	var startTime = time.Now()

	for i := 0; i < 100000000; i++ {
		amountOfNumbersInProgression = (number - 10 + 1)
		sumFromTenProgression = (10 + number) * amountOfNumbersInProgression / 2
	}

	var endTimeFormula = time.Since(startTime)

	console.Writeln("Посчитано по формуле арифметической прогрессии. Сумма всех чисел от 10 до ", number, " = ",
		sumFromTenProgression, "; время обработки формулы, сек = ", endTimeFormula.Seconds())

	if endTimeCycle > endTimeFormula {
		console.Writeln("Цикл выполнялся дольше арифметической прогрессии")
	} else {
		console.Writeln("Цикл выполнялся быстрее арифметической прогрессии")
	}
}
