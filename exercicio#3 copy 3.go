/*
Os fatores primos de 13195 são 5, 7, 13 e 29.
Qual é o maior fator primo do número 600851475143?
*/

package main

import "fmt"

var slice []int
var slice2 []int
var cont = 1

func main() {
	fatPrimo(13195)
}

func fatPrimo(num int) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			slice = append(slice, i)
		}
	}
	fmt.Println("divisões com resto zero:", slice)

	for i, v := range slice {

		fmt.Println(i, v)

	}
	fmt.Println("Contador: ", cont)
	fmt.Println(slice2)
	fmt.Println(slice)
	total := mult(slice...)
	fmt.Println(total)
}

func mult(x ...int) int {
	mult := 0
	for _, v := range x {
		mult = v * v
	}
	return mult
}
