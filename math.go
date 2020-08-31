package goutils

import (
	"math"
)

func Round(val float64, decimal uint8) float64 {
	move := math.Pow10(int(decimal))
	result := val * move
	result = math.Floor(result + 0.5)
	result = result / move
	return result
}
