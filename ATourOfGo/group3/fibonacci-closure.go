package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b, i := 0, 1, -1

	return func() int {
		i++

		if i < 2 {
			return i
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
