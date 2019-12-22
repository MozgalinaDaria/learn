package main

import (
	"../base/console"
	"strings"
)

//TODO: спросить Артема, как решить задачу иначе (баз строк и hash-map)

func main() {
	numFirst := console.ReadString("Введите первое число: ")
	numSecond := console.ReadString("Введите второе число: ")

	console.Writeln("Повторяющиеся цифры в обоих числах: ")
	console.Writeln(compareStrings(numFirst, numSecond))
}

func compareStrings(a, b string) []string {
	var isContained bool
	var commonStrings = make(map[string]int)
	var uniccommonStrings []string

	lengthA := len(a)
	stringToSplit := getStrings(b)

	for i := 0; i < lengthA; i++ {
		isContained = strings.Contains(a, stringToSplit[i])

		if isContained {
			commonStrings[stringToSplit[i]]++
		}
	}

	for key := range commonStrings {
		uniccommonStrings = append(uniccommonStrings, key)
	}
	return uniccommonStrings
}

func getStrings(a string) []string {
	return strings.Split(a, "")
}
