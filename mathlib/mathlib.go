package mathlib

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
