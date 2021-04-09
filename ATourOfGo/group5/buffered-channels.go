package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// gera erro caso o tamanho do buffer seja menor que a quantidade de valores a serem armazenados
	// gera erro caso o tamanho do buffer seja maior que a quantidade de valores a serem armazenados, e essas posições adicionais não tem valor definido
	// fmt.Println(<-ch)
}
