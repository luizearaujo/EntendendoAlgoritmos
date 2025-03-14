package main

import "fmt"

func maiorElemento(x []int) int {
	if len(x) == 1 {
		return x[0]
	}

	m := maiorElemento(x[1:])
	if m > x[0] {
		return m
	}
	return x[0]
}

func main() {
	lista := []int{1, 2, 3, 4, 5}
	result := maiorElemento(lista)
	fmt.Println(result)
}
