package task

import "testing"

type TestCase struct {
	r1                         Rectangle
	r2                         Rectangle
	ExpectedAreCrossing        bool
	ExpectedIntersectionSquare float64
	ExpectedAreNested          int
}

var tests = []TestCase{
	{
		r1:                         Rectangle{Top: 9, Bottom: 6, Left: 2, Right: 5},
		r2:                         Rectangle{Top: 10, Bottom: 3, Left: 1, Right: 7},
		ExpectedAreCrossing:        true,
		ExpectedIntersectionSquare: 9,
		ExpectedAreNested:          0,
	},
}

func TestAreCrossing(t *testing.T) {
	var actualAreCrossing bool

	for _, current := range tests {
		actualAreCrossing = AreCrossing(current.r1, current.r2)
		if actualAreCrossing != current.ExpectedAreCrossing {
			t.Error("For", current.r1, current.r2, "expected", current.ExpectedAreCrossing, "got", actualAreCrossing)
		}
	}
}

func TestIntersektionSquare(t *testing.T) {
	var actualIntersectionSquare float64

	for _, current := range tests {
		actualIntersectionSquare = IntersectionSquare(current.r1, current.r2)
		if actualIntersectionSquare != current.ExpectedIntersectionSquare {
			t.Error("For", current.r1, current.r2, "expected", current.ExpectedIntersectionSquare, "got", actualIntersectionSquare)
		}
	}
}

func TestAreNested(t *testing.T) {
	var actualAreNested int

	for _, current := range tests {
		actualAreNested = AreNested(current.r1, current.r2)
		if actualAreNested != current.ExpectedAreNested {
			t.Error("For", current.r1, current.r2, "expected", current.ExpectedAreNested, "got", actualAreNested)
		}
	}
}
