package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	//A varivel i tem escopo apenas do bloco da condição e das outras condições
	if i := numeroAleatorio(); i > 5 { // Suporte no Switch
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")
	}
	//fmt.Println(i)
}
