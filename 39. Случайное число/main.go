package main

import (
	"../base/console"
	"../base/rand"
	"fmt"
	"strconv"
)

func main() {
	var randomPeriod int
	var randomnumber float64
	var finalnumber float64

	randomPeriod = rand.RandomInt(1, 3)
	switch randomPeriod {
	case 1:
		randomnumber = float64(rand.RandomInt(-1000, -500)) + float64(rand.RandomInt(1, 100))/1000
	case 2:
		randomnumber = float64(rand.RandomInt(0, 100)) + float64(rand.RandomInt(1, 100))/1000
	case 3:
		randomnumber = float64(rand.RandomInt(300, 400)) + float64(rand.RandomInt(1, 100))/1000
	}

	finalnumber, _ = strconv.ParseFloat(fmt.Sprintf("%f", randomnumber), 3)

	console.Writeln("Сгенерированное трехзначное число с тремя знаками после запятой: ", finalnumber)
}
