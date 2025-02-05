package main

import "fmt"

func main() {
	//           tipo key,tipo valor
	var aprovados map[int]string
	// maps devem ser incializados
	aprovados = make(map[int]string)
	aprovados[1] = "Mateus"
	aprovados[2] = "Pedro"
	fmt.Println(aprovados)
	for code, nome := range aprovados {
		fmt.Println(code, nome)
	}
	fmt.Println(aprovados[1])
	delete(aprovados, 1)
	fmt.Println(aprovados[1])

}
