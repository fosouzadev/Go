package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	// o acesso a propriedade deveria ser assim
	(*p).X = 1e4

	// porém a liguagem permite usar assim
	p.X = 1e9

	fmt.Println(v)
}
