package main

import (
	"../base/console"
	"math"
)

func main() {
	var number, even, odd int
	number = int(math.Abs(console.ReadFloat("Введите число: ")))

	for ; number != 0; number = number / 10 {
		figure := number % 10

		if figure%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	console.Writeln("В числе всего", even, " четных и ", odd, " нечетных цифр")
}
