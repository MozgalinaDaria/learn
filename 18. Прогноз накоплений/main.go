package main

import (
	"../base/console"
	"math"
)

func main() {

	sum := console.ReadFloat("Введите исходную сумма вклада: ")
	percent := console.ReadFloat("Введите процентную ставка вклада: ")
	months := console.ReadFloat("Введите срок вклада в месяцах: ")

	finalSum := sum * math.Pow(1+percent/100, months)
	benefit := finalSum - sum

	console.Writeln("Итоговая сумма на конец периода = ", finalSum)
	console.Writeln("Разница между начальной и итоговой суммой = ", benefit)
}
