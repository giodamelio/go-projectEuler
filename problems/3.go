package main

import (
	"fmt"
	"github.com/giodamelio/go-projectEuler/mathlib"
)

func main() {
	var n float64 = 600851475143
	fmt.Println(mathlib.ArrayMax(mathlib.PrimeFactors(n)))
}
