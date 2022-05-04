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
	somaPrimosAbaixoDe(1000)
}
func somaPrimosAbaixoDe(num uint) {
	var i uint
	var sum uint
	var aux uint
	var aux2 uint
	var h uint
	for i = 2; i <= num; i++ {
		/*
			if (i == 2 || i == 3) || (i%2 != 0 && i%3 != 0) && (i > 5 && i%5 != 0) {
				aux = i
				sum += i
				fmt.Println(aux)
			}
		*/

		if (i == 2 || i == 3 || i == 5 || i == 7 || i == 11 || i == 13) || (i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 && i%11 != 0 && i%13 != 0) {
			aux = i
			for h = 14; h < num; h++ {
				if aux%h != 0 {
					//fmt.Println(aux, "/", h)
					aux2 = 0
					sum += aux2

				}
			}
			fmt.Println(aux)
		}
	}
	fmt.Println("\nSoma dos números primos =", sum)
}
