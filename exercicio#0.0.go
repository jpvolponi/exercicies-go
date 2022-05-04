/*Dada uma matriz, encontre o maior número par em uma matriz*/

package main

import (
	"fmt"
)

func FindBiggestEven(numbers []int) int {

	var biggest int
	for _, v := range numbers { //definindo o maior indice da slice
		if v%2 == 0 {
			if biggest < v {
				biggest = v
			}
		}
	}

	return biggest
}

func main() {
	result := FindBiggestEven([]int{3, 8, 14, 7, 2, 6, 7, 9, 4, 12, 5, 10, 11, 0, 13, 17, 19})
	fmt.Println(result)
}
