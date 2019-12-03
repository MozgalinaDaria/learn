package main

import "../base/console"

func main() {
	var lengthOfWord int

	numberOfStringsOnPage := console.ReadInt("Введите количество строк на странице: ")
	numberOfSymbolsInStrings := console.ReadInt("Введите количество символов в строке: ")
	numberOfWords := console.ReadInt("Введите общее количество слов в тексте: ")

	console.Writeln("Введите текст, отделяя каждое слово переносом строки: ")

	balanceOfSymbolsInString := numberOfSymbolsInStrings
	balanceOfStrings := numberOfStringsOnPage
	list := 1

	for i := 0; i < numberOfWords; i++ {

		lengthOfWord = len(console.ReadString())

		if balanceOfStrings > 0 {

			if lengthOfWord < balanceOfSymbolsInString {
				balanceOfSymbolsInString = balanceOfSymbolsInString - lengthOfWord - 1
			} else if lengthOfWord == balanceOfSymbolsInString {
				balanceOfSymbolsInString = 0
			} else if lengthOfWord > balanceOfSymbolsInString {
				balanceOfStrings -= 1
				balanceOfSymbolsInString = numberOfSymbolsInStrings
			}
		} else {
			list += 1
			balanceOfStrings = numberOfStringsOnPage
		}
	}
	console.Writeln("Текст займет листов: ", list)
}
