package main

import (
	"fmt"
	"github.com/giodamelio/go-projectEuler/mathlib"
)

func main() {
	var a, b, largest float64
	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			var c float64 = a * b
			if mathlib.IsPalindrome(c) && c > largest {
				largest = c
			}
		}
	}
	fmt.Println(largest)
}
