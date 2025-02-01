package main

import (
	"fmt"
)

/*
	O slice não é um array, slice define um pedaço do array
	ou seja podemos ter um array base com 20 posições e podemos ter 2 slices
	que apontam pra regigões diferente do array

	Ele não gera um novo array ele simplesmente cria um ponteiro do
	1° indice da slice até fim do slice

	podemos imaginar o slice como: tamanho e um ptr pra um elemento de um array
	podemos fazer slice de outros slices
*/

func main() {

	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	//fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [...]int{1, 2, 3, 4, 5}

	// indices 1,2
	s2 := a2[1:3]
	fmt.Println(a2, s2) // novo slice que aponta do indice 1 até o indice 3

	//indice 0,1,2
	s3 := a2[:3] // novo slice do indice 0 até o indice 3
	fmt.Println(a2, s3)

	//indices 0
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
