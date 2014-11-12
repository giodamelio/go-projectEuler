package main

import (
	"fmt"
)

func fib(limit int) chan int {
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

func main() {
	sum := 0
	for i := range fib(4000000) {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
