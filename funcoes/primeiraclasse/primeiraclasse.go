package main

import "fmt"

//	Tradução de código
//
// Eu estou criando um função anonima e atribuindo a variavel soma
var soma = func(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(5, 10))

}
