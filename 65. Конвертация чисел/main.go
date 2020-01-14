package main

import (
	"../base/console"
	"strings"
)

func main() {
	number := console.ReadString("Введите римское число для конвертации: ")
	console.Writeln("Это соответствует арабскому числу: ", getArabicNumberFromLatinInt(convertLatinStrToInt(number)))
}

func convertLatinStrToInt(number string) []int {
	var ArabicNumber [] int
	output := strings.Split(number, "")
	length := len(output)

	for i := 0; i < length; i++ {
		switch output[i] {
		case "I":
			ArabicNumber = append(ArabicNumber, 1)
		case "V":
			ArabicNumber = append(ArabicNumber, 5)
		case "X":
			ArabicNumber = append(ArabicNumber, 10)
		case "L":
			ArabicNumber = append(ArabicNumber, 50)
		case "C":
			ArabicNumber = append(ArabicNumber, 100)
		case "D":
			ArabicNumber = append(ArabicNumber, 500)
		case "M":
			ArabicNumber = append(ArabicNumber, 1000)
		}
	}
	console.Writeln(ArabicNumber)
	return ArabicNumber
}

func getArabicNumberFromLatinInt(latinInt []int) int {
	var rightNumber, leftNumber, totalArabNumber, newNumber int

	length := len(latinInt) - 1

	for i := length; i > 0; i-- {
		rightNumber = latinInt[i]
		leftNumber = latinInt [i-1]
		if rightNumber <= leftNumber {
			newNumber = leftNumber
		} else {
			newNumber = -leftNumber
		}
		totalArabNumber += newNumber
	}
	return totalArabNumber + latinInt[length]
}
