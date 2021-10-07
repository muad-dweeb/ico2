package commands

import (
	"math"
	"math/rand"
)

func highVariancePercent() {
	tens := []int{
		00,
		00,
		10,
		10,
		20,
		70,
		80,
		80,
		90,
		90,
	}
	ones := []int{
		1,
		1,
		2,
		2,
		3,
		8,
		9,
		9,
		10,
		10,
	}
	result := tens{math.Floor(rand.Int() * len(tens))}
	result += ones{math.Floor(rand.Int() * len(ones))}
	return result
}
