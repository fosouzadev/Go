package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 11, 13}

	s := a[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	b := a[:]
	fmt.Println(b)
}
