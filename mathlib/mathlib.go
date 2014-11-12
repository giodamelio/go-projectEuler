package mathlib

import (
	"math"
	"reflect"
	"strconv"
)

// Get all the fibanacci numbers below a limit
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

// Check to see if a number is a prime
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

// Get the largest element in an array
func ArrayMax(n []float64) float64 {
	var max float64 = 0
	for _, i := range n {
		if i > max {
			max = i
		}
	}
	return max
}

// Get the prime factors of a number
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

// Check to see if the number is a palindrome
func IsPalindrome(n float64) bool {
	reversed := []rune(strconv.FormatInt(int64(n), 10))
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reflect.DeepEqual([]rune(strconv.FormatInt(int64(n), 10)), reversed)
}

// Find the greatest common divisor of two numbers
func GreatestCommonDivisor(a float64, b float64) float64 {
	if b != 0 {
		return GreatestCommonDivisor(b, math.Mod(a, b))
	} else {
		return a
	}
}

// Find the least common multiple of two numbers
func LeastCommonMultiple(a float64, b float64) float64 {
	return (a * b) / GreatestCommonDivisor(a, b)
}

// Get the first n primes
func FirstNPrimes(n float64) []float64 {
	var primes []float64
	var a float64 = 2
	for float64(len(primes)) < n {
		if IsPrime(a) {
			primes = append(primes, a)
		}
		a++
	}
	return primes
}
