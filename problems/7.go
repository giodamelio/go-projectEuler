package main

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10001st prime number?
*/

import (
	"fmt"
	"github.com/giodamelio/go-projectEuler/mathlib"
)

func main() {
	first10001Primes := mathlib.FirstNPrimes(10001)
	fmt.Println(first10001Primes[len(first10001Primes)-1])
}
