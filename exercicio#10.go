/*
A soma dos primos abaixo de 10 é 2 + 3 + 5 + 7 = 17.
Encontre a soma de todos os primos abaixo de dois milhões
*/
package main

import "fmt"

func main() {
	somaPrimosAbaixoDe(100)
}
func somaPrimosAbaixoDe(num uint) {
	var i uint
	var sum uint
	var aux uint
	for i = 2; i <= num; i++ {
		if (i == 2 || i == 3 || i == 4) || (i%2 != 0 && i%3 != 0) {
			aux = i
			sum += i
			fmt.Println(aux)
		}
	}
	fmt.Println("Soma dos números primos =", sum)
}
