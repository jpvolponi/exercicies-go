/*Given an array, sort it in a crescent order*/

package main

import (
	"fmt"
)

func Sort(numbers []int) []int {
	var cont int = 1
	for cont < len(numbers)-1 {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				aux := numbers[i]
				numbers[i] = numbers[i+1]
				numbers[i+1] = aux

			}
		}
		cont++
	}

	return numbers
}

func main() {
	result := Sort([]int{3, 8, 14, 7, 2, 6, 7, 9, 4, 12, 5, 10, 11, 0, 13, 17, 19})
	fmt.Println(result)
}
