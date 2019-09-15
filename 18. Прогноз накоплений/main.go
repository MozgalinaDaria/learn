package main

import (
	"../base/console"
	"math"
)

func main() {
	var sum, percent, months, finalSum, benefit float64

	console.Writeln("Введите исходную сумма вклада, процентную ставка вклада, срок вклада в месяцах")
	sum = console.ReadFloat()
	percent = console.ReadFloat()
	months = console.ReadFloat()

	finalSum = sum * math.Pow(1+percent/100, months)
	benefit = finalSum - sum

	console.Writeln("Итоговая сумма на конец периода = ", finalSum)
	console.Writeln("Разница между начальной и итоговой суммой = ", benefit)
}
