package main

import "fmt"

func buscaBinaria(lista []int, item int) int {
	baixo := 0
	alto := len(lista) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]

		if chute == item {
			return meio
		} else if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1 // Retorna -1 se o item não for encontrado
}

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	item := 2
	resultado := buscaBinaria(lista, item)
	fmt.Printf("Elemento encontrado no índice: %d\n", resultado)
}
