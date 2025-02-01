package main

import "fmt"

func main() {
	//	Homogênea (mesmo tipo) e estática (tamanho fixo)
	// O array é um estrutura indexavel inicinaod com indice 0
	var notas [3]float64

	//fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 6.2, 9.5
	//notas[3] = 10 // Erro pois o index 3 não existe no array
	var total float64
	for i := 0; i < len(notas); i++ {
		total = notas[i]
	}
	//	como total é do tipo float64 pra poder relizar a divisão do total
	// pela media eu preciso fazer o casting de tipo do tamanho do array
	media := total / float64(len(notas))
	//fmt.Println(notas)
	fmt.Printf("A media das notas é %.2f\n", media)
}
