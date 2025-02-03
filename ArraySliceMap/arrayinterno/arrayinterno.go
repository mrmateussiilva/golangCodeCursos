package main

import "fmt"

func main() {
	slice1 := make([]int, 10, 15)
	slice2 := append(slice1, 10, 2, 3)

	fmt.Println(slice1, slice2)
	slice1[1] = 255
	fmt.Println(slice1, slice2)

}
