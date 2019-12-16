package main

import "../base/console"

// TODO Вынести определение четверти в отдельную функцию func DetectPosition(x, y int) int
// TODO Функция должна возвращать константы (своя константа для каждой ситуации)
// TODO func должна определять что вывести пользователю через switch/case

const BetweenSecondAndFirstQuarters = 1
const BetweenThirdAndFourthQuarters = 2
const InFirstQuarter = 3
const BetweenFirstAndFourthQuarters = 4
const InFourthQuarter = 5
const InThirdQuarter = 6
const BetweenSecondAndThirdQuarters = 7
const InSecondQuarter = 8
const OnBorderOfAllQuarters = 9

func main() {
	x := console.ReadInt("Введите координаты точки по оси Х: ")
	y := console.ReadInt("Введите координаты точки по оси Y: ")

	switch DetectPosition(x, y) {
	case BetweenSecondAndFirstQuarters:
		console.Writeln("Координата находится на границе 2-й и 1-ой четвертей")
	case BetweenThirdAndFourthQuarters:
		console.Writeln("Координата находится на границе 3-ей и 4-ой четвертей")
	case InFirstQuarter:
		console.Writeln("Координата находится в 1-й четверти")
	case BetweenFirstAndFourthQuarters:
		console.Writeln("Координата находится на границе 1-й и 4-й четвертей")
	case InFourthQuarter:
		console.Writeln("Координата находится в 4-й четверти")
	case InThirdQuarter:
		console.Writeln("Координата находится в 3-ей четвертей")
	case BetweenSecondAndThirdQuarters:
		console.Writeln("Координата находится на границе 2-й и 3-й четвертей")
	case InSecondQuarter:
		console.Writeln("Координата находится во 2-й четверти")
	case OnBorderOfAllQuarters:
		console.Writeln("Координата находится на границе всех четырех четвертей")
	}
}

func DetectPosition(x, y int) int {
	if x == 0 {
		if y > 0 {
			return BetweenSecondAndFirstQuarters
		}

		if y < 0 {
			return BetweenThirdAndFourthQuarters
		}
	}

	if x > 0 {
		if y > 0 {
			return InFirstQuarter
		}

		if y == 0 {
			return BetweenFirstAndFourthQuarters
		}

		return InFourthQuarter
	}

	if x < 0 {
		if y < 0 {
			return InThirdQuarter
		}

		if y == 0 {
			return BetweenSecondAndThirdQuarters
		}

		return InSecondQuarter
	}

	return OnBorderOfAllQuarters
}
