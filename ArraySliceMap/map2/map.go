package main

import "fmt"

func main() {

	// Declarando e inicializando o map de forma literal
	funcaSalrio := map[string]float64{
		//key :  value
		"Mateus": 1110.00,
		"Jose":   15.00,
		"Silva":  9800.00,
	}
	funcaSalrio["Rafel"] = 1350.00
	delete(funcaSalrio, "Inexistente")
	for nome, salario := range funcaSalrio {
		fmt.Printf("%s R$%.2f\n", nome, salario)
	}
}
