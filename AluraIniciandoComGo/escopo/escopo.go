package main

import "fmt"

func main() {
	i := 1
	fmt.Println(i)

	{
		i := 2
		i++
		fmt.Println(i)
	}
}
