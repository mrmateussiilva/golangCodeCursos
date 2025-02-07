package main

import "fmt"

func main() {

	//                   CHAVE VALOR CHAVE VALOR
	funcsPorLetra := map[string]map[string]float64{
		"A": {
			"Amanda":   1556.00,
			"Anderson": 9001.00,
		},
		"M": {
			"Mateus": 199.80,
			"Maicon": 9888.00,
		},
	}
	//delete(funcsPorLetra, "A") // DELETA O MAP

	for letra := range funcsPorLetra {
		fmt.Println(letra)
		for k, v := range funcsPorLetra[letra] {
			fmt.Println("\t", k, v)
		}
		fmt.Println(letra)
	}
	//var funcs map[string]map[string]string // funciona assim

}
