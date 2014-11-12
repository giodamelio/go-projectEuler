package main

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

import (
	"fmt"
	"github.com/giodamelio/go-projectEuler/mathlib"
)

func main() {
	var a, b, c, tripletProduct float64
	for a = 1; a <= 1000; a++ {
		for b = 1; b <= (1000 - a); b++ {
			for c = 1; c <= (1000 - b); c++ {
				if a+b+c == 1000 {
					if mathlib.IsPythagoreanTriplet(a, b, c) {
						tripletProduct = a * b * c
						break
					}
				}
			}
		}
	}
	fmt.Printf("%f\n", tripletProduct)
}
