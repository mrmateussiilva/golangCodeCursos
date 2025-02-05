package main

import "fmt"

func main() {
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

	for letra := range funcsPorLetra {
		fmt.Println(letra)
		for k, v := range funcsPorLetra[letra] {
			fmt.Println("\t", k, v)
		}
		//fmt.Println(letra)
	}
	estoqueMercado :=
}
