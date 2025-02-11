package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	if len(numeros) == 0 {
		return 0.0
	}

	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média %.2f\n", media(10, 20, 30, 40))
	fmt.Printf("Média %.2f\n", media())
}
