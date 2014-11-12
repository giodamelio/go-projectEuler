package main

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
