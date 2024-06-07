package mymath

import "math"

func Sqrt(x float64) float64 {
    return math.Sqrt(x) 
}

func Pow(x, y float64) float64 {
    result := 1.0
    for i := 0; i < int(y); i++ {
        result *= x
    }
    return result
}

func Max(x, y float64) float64 {
    if x > y {
        return x
    }
    return y
}
