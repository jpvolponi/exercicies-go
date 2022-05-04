//Given an array, find the lowest even number in an array
package main

import "fmt"

func FindLowestEven(numbers []uint) uint {
	var lowest uint
	var biggest uint
	for _, v := range numbers { //definindo o maior indice da slice
		if v%2 == 0 {
			if biggest < v {
				biggest = v
			}
		}
	}
	lowest = biggest //passando o maior para a variavel do menor
	//return biggest
	for _, v := range numbers {
		if v%2 == 0 {
			if lowest > v { //comparando o maior com os indices do slice para achar o menor
				lowest = v
			}
		}
	}
	return lowest

}

func main() {
	result := FindLowestEven([]uint{3, 8, 14, 7, 2, 6, 7, 9, 4, 12, 5, 10, 11, 0, 13, 17, 19})
	fmt.Println(result)
}
