package main

import (
	"../base/console"
)

func main() {
	number := 0
	list := 0
	j := 0

	numberOfStringsOnPage := console.ReadInt("Введите количество строк на странице: ")
	numberOfSymbolsInStrings := console.ReadInt("Введите количество символов в строке: ")
	numberOfWords := console.ReadInt("Введите общее количество слов в тексте: ")

	console.Writeln("Введите текст, отделяя каждое слово переносом строки:")

	input := make([]string, numberOfWords)

	for i := 0; i < numberOfWords; i++ {
		input[i] = console.ReadString()
	}

	for i := 0; i < numberOfWords; i = number * j * list {
		list += 1

		for j = 0; j < numberOfStringsOnPage; j++ {

			for k := 0; k < numberOfSymbolsInStrings; k = len(input[number]) + 1 {

				number += 1 //количество слов в строке
			}
		}
	}
}