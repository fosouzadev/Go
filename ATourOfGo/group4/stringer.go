package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string
	Age  int
}

// implementação da interface Stringer do package fmt
// equivalente ao ToString()
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	// tipo que não implementa Stringer
	y := Person2{Name: "Felipe", Age: 31}

	fmt.Println(a, z, y)
}
