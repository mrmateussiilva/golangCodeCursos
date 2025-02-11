package main

import "fmt"

/*
Em go temos o retorno nomeado que é como o nome diz, nomeamos o retorno da função

*/

func trocar(p1, p2 int) (segundo, primeiro int) {
	// segundo = p2
	// primeiro = p1
	return p2, p1
	// return // retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
