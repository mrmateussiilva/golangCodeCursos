package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // Data e hora atual
	//fmt.Println(t)
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Bom tarde")
	default:
		fmt.Println("Boa Noite")
	}

}
