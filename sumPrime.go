package main

import (
	"fmt"
	_ "fmt"
)

func sumPrime(x uint) uint {

	var i uint

	var sum uint
	for i = 2; i < x; i++ {
		var isDiv bool
		if x%i == 0 {
			isDiv = true
		}
		if isDiv == false {

			sum += i
		}
		_ = isDiv
	}
	return sum
}

func main() {
	fmt.Println(sumPrime(10))
}
