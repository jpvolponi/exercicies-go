/*
A soma dos primos abaixo de 10 é 2 + 3 + 5 + 7 = 17.
Encontre a soma de todos os primos abaixo de dois milhões

Regras para um número primo:

- Positivos;
- Divisíveis por 1 e por eles mesmos somente;
- Números pares não são primos, com exceção do número 2;
- Não existem números primos terminados em 5, com exceção do próprio 5;
-

*/
package main

import "fmt"

func main() {
	fmt.Println(somaPrimosAbaixoDe(1000))
}
func somaPrimosAbaixoDe(num uint) uint {
	var i uint
	var sum uint
	//var aux uint
	for i = 2; i <= num; i++ {
		if divisivelPor(i) == false {
			//aux = i
			sum += i
			//fmt.Println(aux)
		}
	}
	return sum
}

func divisivelPor(x uint) bool {
	var isDiv bool
	var i uint
	for i = 2; i < x; i++ {
		if x%i == 0 {
			isDiv = true
		}
	}
	return isDiv
}
