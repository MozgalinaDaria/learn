package main

import "../base/console"

const IsOdd int = 1
const IsEven int = 2

func main() {
	number := console.ReadInt("Введите число для проверки гипотезы Сиракуз: ")
	SyracuseHypothesis(number)
}

func SyracuseHypothesis(number int) {
	for number > 1 {
		switch CheckParity(number) {
		case IsEven:
			number /= 2
			console.Write("|", number, "|", " ")
		case IsOdd:
			number = (number*3 + 1) / 2
			console.Write("|", number, "|", " ")
		}

	}
}

func CheckParity(a int) int {
	if a%2 == 0 {
		return 2
	}
	return 1
}
