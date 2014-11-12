package main

import (
	"fmt"
	"github.com/giodamelio/go-projectEuler/mathlib"
)

func main() {
	sum := 0
	for i := range mathlib.Fib(4000000) {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
