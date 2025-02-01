package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	case bool:
		return "Booleano"
	default:
		return "Não conheço esse tipo"
	}

}

func main() {
	fmt.Println(tipo(time.Now()))
	fmt.Println(tipo(42))
	fmt.Println(tipo(3.14))
	fmt.Println(tipo("Mateus"))
	fmt.Println(tipo(main))
	fmt.Println(tipo(false))
}
