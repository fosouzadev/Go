package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// a function pos e neg estão vinculadas a uma variavel sum
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {

		// a cada chamada de pos ou neg, sum é alterada e mantem seu valor para uma proxima chamada
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
