package task


import "math"

type Rectangle struct{
	Left float64
	Right float64
	Top float64
	Bottom float64
}

func AreCrossing(r1, r2 Rectangle) bool {
	if (r1.Top < r2.Bottom || r2.Bottom < r1.Bottom) && (r1.Left < r2.Right || r2.Right < r1.Right){
		return true
	}

	if (r2.Top < r1.Bottom || r1.Bottom < r2.Bottom) && (r2.Left < r1.Right || r1.Right < r2.Right){
		return true
	}

	return false
}

func IntersectionSquare(r1, r2 Rectangle) float64 {
	var OX, OY float64

	OX = math.Max(r1.Left, r2.Left) - math.Min(r1.Right, r2.Right)
	OY = math.Min(r1.Top, r2.Top) - math.Max(r1.Bottom, r2.Bottom)

	return math.Abs(OX * OY)
}

func AreNested(r1, r2 Rectangle) int {
	if r1.Top < r2.Top && r1.Bottom > r2.Bottom && r1.Left > r2.Left && r1.Right < r2.Right {
		return 0
	}

	if r2.Top < r1.Top && r2.Bottom > r1.Bottom && r2.Left > r1.Left && r2.Right < r1.Right {
		return 1
	} else {
		return -1
	}
}

