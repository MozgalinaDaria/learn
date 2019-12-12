package main

import (
	"../base/console"
	"fmt"
)

//TODO Mozgalina: задача решена неверно. Делиться только сами на себя помимо 11 и тп может, например, 37. Поэтому решить задачу полным перебором

func main() {
	var min, max, numberToFindDividers int

	fmt.Println("Введите 2 числа, с какого по какое вывести делители: ")
	fmt.Scan(&min, &max)

	for i := min; i <= max; i++ {
		numberToFindDividers = i
		console.Writeln("Делители для ", numberToFindDividers, " = ", getDividers(numberToFindDividers))
	}
}

func getDividers(numberToFindDividers int) []int {
	var secondDivider int
	var keys [] int
	var specialDividers = [4] int{11, 13, 17, 19}

	dividers := make(map[int]int)

	for j := 2; j <= 10; j++ {
		if numberToFindDividers%j == 0 {
			secondDivider = numberToFindDividers / j
			dividers[j]++
			dividers[secondDivider]++
		}
	}

	for i := 0; i < 4; i++ {
		if numberToFindDividers%specialDividers[i] == 0 {
			secondDivider = numberToFindDividers / specialDividers[i]
			dividers[specialDividers[i]]++
			dividers[secondDivider]++
		}
	}

	for key := range dividers {
		keys = append(keys, key)
	}

	return keys
}
