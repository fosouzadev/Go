package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	t := make(map[string]Vertex)

	t["Answer"] = Vertex{1, 2}
	fmt.Println("The value:", t["Answer"])

	t["Answer"] = Vertex{3, 4}
	fmt.Println("The value:", t["Answer"])

	delete(t, "Answer")
	fmt.Println("The value:", t["Answer"])

	tv, tok := t["Answer"]
	fmt.Println("The value:", tv, "Present?", tok)
}
