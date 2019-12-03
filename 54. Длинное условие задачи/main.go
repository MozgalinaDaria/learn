package main

import "../base/console"
import "../base/strings"

type Writable interface {
	CanWrite(word string) bool
	Write(word string)
}

// -------------------------------

type Line struct {
	Writable

	CurrentSymbol int
	SymbolsLimit  int
}

func (object *Line) CanWrite(word string) bool {
	return object.CurrentSymbol+len(word) <= object.SymbolsLimit
}

func (object *Line) Write(word string) {
	if !object.CanWrite(word) {
		panic("Can't write word. Line will be overload.")
	}

	object.CurrentSymbol += len(word) + 1
}

func NewLine(SymbolsLimit int) *Line {
	return &Line{
		CurrentSymbol: 0,
		SymbolsLimit:  SymbolsLimit,
	}
}

// -------------------------------

type Page struct {
	Writable

	CurrentLine *Line
	LinesCount  int

	SymbolsLimit int
	LinesLimit   int
}

func (object *Page) Write(word string) {
	if !object.CanWrite(word) {
		panic("Can't write word. Page will be overload.")
	}

	if object.CurrentLine.CanWrite(word) {
		object.CurrentLine.Write(word)
		return
	}

	object.LinesCount += 1
	object.CurrentLine = NewLine(object.SymbolsLimit)
	object.CurrentLine.Write(word)
}

func (object *Page) CanWrite(word string) bool {
	if object.CurrentLine.CanWrite(word) {
		return true
	}

	return object.LinesCount < object.LinesLimit
}

func NewPage(SymbolsLimit, LineLimit int) *Page {
	return &Page{
		CurrentLine:  NewLine(SymbolsLimit),
		LinesCount:   1,
		SymbolsLimit: SymbolsLimit,
		LinesLimit:   LineLimit,
	}
}

// -------------------------------

type PagesCounter struct {
	SymbolsLimit int
	LinesLimit   int
	PagesCount   int
	CurrentPage  *Page
}

func (object *PagesCounter) WriteWord(word string) {
	if object.CurrentPage != nil && object.CurrentPage.CanWrite(word) {
		object.CurrentPage.Write(word)
		return
	}

	object.PagesCount += 1
	object.CurrentPage = NewPage(object.SymbolsLimit, object.LinesLimit)
	object.CurrentPage.Write(word)
}

func (object *PagesCounter) GetPagesCount() int {
	return object.PagesCount
}

func NewPagesCounter(SymbolsLimit, LinesLimit int) *PagesCounter {
	return &PagesCounter{
		LinesLimit:   LinesLimit,
		SymbolsLimit: SymbolsLimit,
		PagesCount:   0,
		CurrentPage:  nil,
	}
}

// -------------------------------

func main() {
	numberOfLinesOnPage := console.ReadInt("Введите количество строк на странице: ")
	numberOfSymbolsInLine := console.ReadInt("Введите количество символов в строке: ")
	numberOfWords := console.ReadInt("Введите общее количество слов в тексте: ")

	console.Writeln("Введите текст, отделяя каждое слово переносом строки: ")

	counter := NewPagesCounter(numberOfSymbolsInLine, numberOfLinesOnPage)

	for i := 0; i < numberOfWords; i++ {
		counter.WriteWord(console.ReadString())
	}

	count := counter.GetPagesCount()
	pageTitle := strings.Pluralize(count, "лист", "листа", "листов")
	console.Writeln("Текст займет", count, pageTitle)
}
