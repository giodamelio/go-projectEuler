package mathlib

import (
	"math"
)

func Fib(limit int) chan int {
	c := make(chan int)

	go func() {
		a, b := 0, 1

		for a < limit {
			c <- a
			a, b = b, a+b
		}

		close(c)
	}()

	return c
}

func IsPrime(n float64) bool {
	sqrtN := math.Sqrt(n)
	var i float64
	for i = 2; i <= sqrtN; i++ {
		if math.Mod(n, i) == 0 {
			return false
		}
	}
	return true
}

func ArrayMax(n []float64) float64 {
	var max float64 = 0
	for _, i := range n {
		if i > max {
			max = i
		}
	}
	return max
}

func PrimeFactors(n float64) []float64 {
	var primeFactors []float64
	var i float64
	for i = 2; i <= math.Sqrt(n); i++ {
		if math.Mod(n, i) == 0 && IsPrime(i) {
			primeFactors = append(primeFactors, i)
		}
	}
	return primeFactors
}
