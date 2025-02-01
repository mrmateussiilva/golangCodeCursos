package main

import "fmt"

func main() {
	/*
		make([]tipo,tamanho,capacidade)
		O make internamente vai trabalhar com array

	*/
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Printf("Slice %d, \ntamanho: %d, \ncapicidade %d\n", s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 15)
	fmt.Println(s, len(s), cap(s))

}
