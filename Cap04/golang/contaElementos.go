package main

import "fmt"

func contaElementos(x []int) int {
	if len(x) == 1 {
		return 1
	}
	return 1 + contaElementos(x[1:])
}

func main() {
	lista := []int{1, 2, 3, 4, 5}
	result := contaElementos(lista)
	fmt.Println(result)
}
