package main

import "fmt"

/*
Para definir uma função em go é bastante simples

func identificador(p)


*/

// Não recebe e não retorna nada
func f1() {
	fmt.Println("F1")
}

// Recene 2 parametros mas não retorna nada
func f2(p1 string, p2 string) {
	fmt.Printf("FS: %s %s\n", p1, p2)
}

// Se prometeu que vai retornar valor precisa retornar
// go vai te exigir isso
// Não recebe para metro mas retorna uma string
func f3() string {
	return "F3"
}

// se os paraementros forem dos mesmo tipo
// você pode simplificar assim
// Recebe 2 parametros e retorna 1 valor
func f4(p1, p2 string) string {
	return fmt.Sprintf("%s %s", p1, p2)
}

// Não recebe parametros e retorna dois valores
func f5() (error, string, int) {
	return nil, "F5", 9
}

func main() {
	f1()
	f2("Mateus", "jose")
	r3, r4 := f3(), f4("Parametro 1", "parametros 2")
	fmt.Println(r3)
	fmt.Println(r4)

	err, r5, r6 := f5()
	if err == nil {
		fmt.Println(r5)
		fmt.Println(r6)
	}
}
