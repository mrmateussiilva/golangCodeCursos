package main

import "fmt"

func notParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // SE der match entra aqui mas não encera o switch
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Invalida"
	}
}

func main() {
	var notas = [...]int{1, 3, 6, 9, 7, 24, 9}
	for i := range notas {
		res := notParaConceito(float64(notas[i]))
		fmt.Printf("A nota %d em conceito é %s\n", notas[i], res)

	}
}
