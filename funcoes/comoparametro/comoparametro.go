package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	result := exec(mult, 2, 10)
	fmt.Println(result)
}
