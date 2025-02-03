package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int
	//Adiciona mas elementos a um slice,
	// a função append retorna um slice novo
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	/*A função copy copia dados de um slice pra outro slice
	caso o tamanho do slice de destino seja menor em comparação com a origem
	o metodo só vai copiar os elementos que cabe
	ele não expande o slice apenas copia os elementos de um lador pra outro
	*/

	//  DESTINO, ORIGEM
	copy(slice2, slice1)

	fmt.Println(slice2)

	s := make([]int, 6)
	t := append(s, 4, 5, 6)
	fmt.Println(t)

}
