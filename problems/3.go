package main

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

import (
	"fmt"
	"github.com/giodamelio/go-projectEuler/mathlib"
)

func main() {
	var n float64 = 600851475143
	fmt.Println(mathlib.ArrayMax(mathlib.PrimeFactors(n)))
}
