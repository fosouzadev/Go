package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 0
	var first bool = true
	var second bool = false

	return func() int {
		if first {
			first = false
			second = true
			return 0
		}

		if second {
			second = false
			b = 1
			return b
		}

		res := a + b
		a = b
		b = res

		return res
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
