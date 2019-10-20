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

	orderedSyllables := make(map[string][]string)
	for i := 0; i < numberOfSyllables; i++ {
		firstLetterInSyllable, _ := utf8.DecodeRuneInString(syllables[i])
		orderedSyllables[string(firstLetterInSyllable)]= append(orderedSyllables[string(firstLetterInSyllable)], syllables[i])
		}
	console.Writeln("Строки, начинающиеся с ", letter, ": ")
	console.Writeln(orderedSyllables[letter])
}


//data := make(map[string][]string)
//
//data["a"] = append(data["a"], "abc")
//data["a"] = append(data["a"], "abd")
//data["a"] = append(data["a"], "abk")
//
//data["b"] = append(data["b"], "bbd")
