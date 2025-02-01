package main

import "fmt"

func main() {
	/*Como estamos usandos a syntax de ... pra falar o tamanho do array
	durante o processo de compilação o compilador vai contar quantos elementos
	existe nesse array
	*/
	// OBS: Se esquecer esses ... go vai tratar isso como um slice e não como array

	numeros := [...]int{1, 2, 3, 4, 5} // size 5
	//Percorrendo o array com for range
	fmt.Println("Pegando tanto o index quanto o value")
	for index, value := range numeros {
		fmt.Printf("Posição %d, valor %d\n", index, value)
	}
	/*Quando usamos o _(underline) assim dizemos pro go,
	"ingora esse valor ou pacote pois não uso ele

	*/
	fmt.Println("Pegando apenas o valor e ignorando o index")
	for _, value := range numeros {
		fmt.Println(value)
	}

	// Tbm podemos pegar apenas o indice da estrutura
	fmt.Println("Pegando apenas o index")
	for index := range numeros {
		fmt.Println(index)
	}
}
