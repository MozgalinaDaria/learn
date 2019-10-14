package main

import (
"../base/console"
"../base/rand"
	"time"
)

func main() {
	var sumFromTenCycle, sumFromTenProgression, amountOfNumbersInProgression int

	number := rand.RandomInt(50, 100)
	console.Writeln("Рандомно было сгенерировано число: ", number)

	var startTime = time.Now()
	for i := 10; i <= number; i++ {
		sumFromTenCycle += i
	}
	var endTime  =  time.Since(startTime)

	console.Writeln("Посчитано в цикле. Сумма всех чисел от 10 до ", number, " = ", sumFromTenCycle, "; время обработки цикла = ", endTime.Nanoseconds())

	startTime = time.Now()
	amountOfNumbersInProgression = (number - 10 + 1)
	sumFromTenProgression = (10 + number) * amountOfNumbersInProgression / 2
	endTime  =  time.Since(startTime)
	console.Writeln("Посчитано по формуле арифметической прогрессии. Сумма всех чисел от 10 до ", number, " = ", sumFromTenProgression , "; время обработки цикла = ", endTime.Nanoseconds())
}

