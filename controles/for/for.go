package main

import (
	"fmt"
	"time"
)

func main() {
	var c int = 0
	for c <= 10 { // não tem while em go
		fmt.Println(c)
		c++
	}
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nMisturando loop com condicional....")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par")
		} else {
			fmt.Print("Impar")
		}
	}

	//	Lopp infinito
	for {
		fmt.Println("\nEsse código é infinito...")
		time.Sleep(time.Second)
		break
	}
}
