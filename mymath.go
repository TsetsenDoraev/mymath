package mymath

import "math"

func Sqrt(x float64) float64 {
    return math.Sqrt(x) 
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}
