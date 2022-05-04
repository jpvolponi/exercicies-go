//Given an array, sum all evens numbers in an array

package main

import "fmt"

func SumEvens(numbers []uint) uint {
	var sum uint
	for _, v := range numbers {
		if v%2 == 0 {
			sum += v
		}
	}
	return sum
}

func main() {
	result := SumEvens([]uint{3, 8, 14, 7, 2, 6, 7, 9, 4, 12, 5, 10, 11, 0, 13, 17, 19})
	fmt.Println(result)

}
