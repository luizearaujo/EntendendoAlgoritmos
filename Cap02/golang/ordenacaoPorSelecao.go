package main

import "fmt"

func buscaMenor(arr []int) int {
	menor := arr[0]
	menor_indice := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < menor {
			menor = arr[i]
			menor_indice = i
		}
	}

	return menor_indice
}

func ordenacaoPorSelecao(arr []int) []int {
	novoArr := []int{}
	for len(arr) > 0 {
		menor := buscaMenor(arr)
		novoArr = append(novoArr, arr[menor])
		arr = append(arr[:menor], arr[menor+1:]...)
	}
	return novoArr
}

func main() {
	lista := []int{5, 3, 6, 2, 10}
	novaLista := ordenacaoPorSelecao(lista)
	for i := 0; i < len(novaLista); i++ {
		fmt.Printf("%d, ", novaLista[i])
	}
}
