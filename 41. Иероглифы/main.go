package main

import (
	"../base/console"
	"unicode/utf8"
)

func main(){
	var numberOfSyllables int
	var syllables []string
	var letter string

	numberOfSyllables = console.ReadInt("Введите количество слогов: ")
	
	for i := 0; i < numberOfSyllables; i++ {
		syllables = append(syllables, console.ReadString("Введите слог: "))
	}

	letter = console.ReadString("Введите букву, на которую искать: ")

	console.Writeln("Строки, начинающиеся с ", letter, ": ")

	for i := 0; i < numberOfSyllables; i++ {
		firstLetterInSyllable, _ := utf8.DecodeRuneInString(syllables[i])
		if letter == string(firstLetterInSyllable) {
			console.Writeln(syllables[i])
		}
	}
}
