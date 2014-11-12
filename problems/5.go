package main

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

import (
	"fmt"
	"github.com/giodamelio/go-projectEuler/mathlib"
)

func main() {
	var max float64 = 20
	var smallest float64 = 1
	var i float64
	for i = 1; i <= max; i++ {
		smallest = mathlib.LeastCommonMultiple(smallest, i)
	}
	fmt.Printf("%f\n", smallest)
}
