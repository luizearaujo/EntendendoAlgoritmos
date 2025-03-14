package main

import "fmt"

func soma(x []int) int {
	if len(x) == 0 {
		return 0
	}
	if len(x) == 1 {
		return x[0]
	}

	return x[0] + soma(x[1:])
}

func main() {
	lista := []int{1, 2, 3, 4, 5}
	result := soma(lista)
	fmt.Println(result)
}
