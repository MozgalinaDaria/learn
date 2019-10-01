package task

import "testing"

type TestCase struct {
	Number           int
	Nominative       string
	GenitiveSingular string
	GenitivePlural   string
	Expected         string
}

var tests = []TestCase{
	{Number: 1, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "заяц"},
	{Number: 2, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайца"},
	{Number: 3, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайца"},
	{Number: 4, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайца"},
	{Number: 5, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 6, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 7, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 8, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 9, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 10, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 11, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 12, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 13, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 14, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 15, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 16, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 17, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 18, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 19, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 20, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
	{Number: 21, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "заяц"},
	{Number: 22, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайца"},
	{Number: 101, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "заяц"},
	{Number: 99912, Nominative: "заяц", GenitiveSingular: "зайца", GenitivePlural: "зайцев", Expected: "зайцев"},
}

func TestDeclension(t *testing.T) {
	var actual string
	for _, current := range tests {
		actual = Declension(current.Number, current.Nominative, current.GenitiveSingular, current.GenitivePlural)
		if actual != current.Expected {
			t.Error("For", current.Number, "expected", current.Expected, "got", actual)
		}
	}
}
