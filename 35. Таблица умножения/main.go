package main

import (
	"../base/console"
	"fmt"
	"strconv"
)

func main() {
	var maxNumber, i, j, numSize, resultSize int
	var format, nsString, rsString string
	maxNumber = console.ReadInt("Введите число, до которого вывести таблицу умножения: ")

	numSize = len(strconv.Itoa(maxNumber))
	resultSize = len(strconv.Itoa(maxNumber * maxNumber))
	nsString = strconv.Itoa(numSize)
	rsString = strconv.Itoa(resultSize)

	for i = 2; i <= maxNumber; i++ {
		for j = 2; j <= maxNumber; j++ {
			format = "%" + nsString + "d * %" + nsString + "d = %" + rsString + "d\n"
			fmt.Printf(format, i, j, i*j)
		}
	}
}
