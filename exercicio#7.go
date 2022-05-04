/*
“By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?”

*/

package main

import "fmt"

func findPrime(x uint) bool {
	var isDiv bool
	var i uint
	for i = 2; i < x; i++ {
		if x%i == 0 {
			isDiv = true
		}
	}
	return isDiv
}

func primoPosicao(num uint) uint {
	var cont uint = 0
	var x uint = 2
	var result uint = 0
	for cont < num {
		if findPrime(x) == false {
			cont += 1
			result = x
		}
		x++
	}
	//fmt.Println("Debug:", cont)
	return result
}

func main() {
	fmt.Println(primoPosicao(10_001))
}
