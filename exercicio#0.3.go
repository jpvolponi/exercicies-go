/*Given an array, sort it in a crescent order*/

package main

import (
	"fmt"
	"sort"
)

func Sort(numbers []int) []int {

	sort.Slice(numbers, func(i, j int) bool {

		return numbers[i] < numbers[j]

	})

	return numbers
}

func main() {
	result := Sort([]int{3, 8, 14, 7, 2, 6, 7, 9, 4, 12, 5, 10, 11, 0, 13, 17, 19})
	fmt.Println(result)
}
