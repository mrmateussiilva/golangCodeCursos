package main

import "fmt"

func notaParaConceio(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 7 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {

	var notas = [...]int{1, 3, 6, 9, 7, 24, 9}
	for i := range notas {
		res := notaParaConceio(float64(notas[i]))
		fmt.Printf("A nota %d em conceito é %s\n", notas[i], res)

	}
}
