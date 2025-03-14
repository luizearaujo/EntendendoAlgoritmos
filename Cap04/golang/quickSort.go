package main

import "fmt"

func qSort(x []int) []int {
	if len(x) < 2 {
		return x
	} else {
		pivot := x[0]
		less := []int{}
		higher := []int{}
		for i := 1; i < len(x); i++ {
			if x[i] <= pivot {
				less = append(less, x[i])
			} else {
				higher = append(higher, x[i])
			}
		}
		return append(append(qSort((less)), pivot), qSort(higher)...)
	}
}

func main() {
	lista := []int{3, 6, 8, 10, 1, 2, 1}
	result := qSort(lista)
	fmt.Println(result)
}
